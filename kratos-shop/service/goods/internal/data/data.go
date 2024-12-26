package data

import (
	"context"
	"github.com/redis/go-redis/extra/redisotel/v9"
	"goods/internal/biz"

	"github.com/redis/go-redis/v9"
	"goods/internal/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	slog "log"
	"os"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData, NewMySQL, NewRedis,
	NewCategoryRepo,
	NewGoodsTypeRepo,
	NewSpecificationRepo,
	NewGoodsAttrRepo,
	NewTransaction,
)

// Data .
type Data struct {
	db  *gorm.DB
	rdb *redis.Client
}

// 用来承载事务的上下文
type contextTxKey struct{}

// NewData .
func NewData(c *conf.Data, db *gorm.DB, rdb *redis.Client, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{
		db:  db,
		rdb: rdb,
	}, cleanup, nil
}

// NewMySQL returns a new gorm db instance.
func NewMySQL(c *conf.Data) *gorm.DB {
	newLogger := logger.New(
		slog.New(os.Stdout, "\r\n", slog.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second, //慢查询SQL阈值
			Colorful:      true,        // 开启彩色打印
			LogLevel:      logger.Info, // 日志级别
		})
	db, err := gorm.Open(mysql.Open(c.Mysql.Source), &gorm.Config{
		Logger:                                   newLogger,
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic(err)
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(int(c.Mysql.MaxIdle))
	sqlDB.SetMaxOpenConns(int(c.Mysql.MaxOpen))
	if err := sqlDB.Ping(); err != nil {
		panic(err)
	}
	tables := []interface{}{
		&Category{},
		&GoodsType{},
		&GoodsAttrValue{},
		&GoodsAttr{},
		&GoodsAttrGroup{},
		&SpecificationsAttr{},
		&SpecificationsAttrValue{},
	}
	err = db.AutoMigrate(tables...)
	if err != nil {
		panic(err)
	}
	return db
}

// NewRedis returns a new redis client instance.
func NewRedis(c *conf.Data) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:         c.Redis.Addr,
		Password:     c.Redis.Password,
		DB:           int(c.Redis.Db),
		DialTimeout:  c.Redis.DialTimeout.AsDuration(),
		ReadTimeout:  c.Redis.ReadTimeout.AsDuration(),
		WriteTimeout: c.Redis.WriteTimeout.AsDuration(),
	})
	// 启用 tracing
	if err := redisotel.InstrumentTracing(rdb); err != nil {
		panic(err)
	}
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
	return rdb
}

// NewTransaction .
func NewTransaction(d *Data) biz.Transaction {
	return d
}

// ExecTx 执行事务
func (d *Data) ExecTx(ctx context.Context, fn func(ctx context.Context) error) error {
	return d.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		ctx = context.WithValue(ctx, contextTxKey{}, tx)
		return fn(ctx)
	})
}

// DB 判断当前db使用的是不是事务的 db
func (d *Data) DB(ctx context.Context) *gorm.DB {
	tx, ok := ctx.Value(contextTxKey{}).(*gorm.DB)
	if ok {
		return tx
	}
	return d.db
}
