package data

import (
	"kratos-realworld/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewMySQL, NewTagRepo, NewUserRepo, NewCommentRepo, NewArticleRepo)

// Data .
type Data struct {
	mysqlDB *gorm.DB
}

// NewData .
func NewData(c *conf.Data, logger log.Logger, mysqlDB *gorm.DB) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{mysqlDB: mysqlDB}, cleanup, nil
}
func NewMySQL(c *conf.Data) *gorm.DB {
	dsn := c.Database.Source
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database:" + err.Error())
	}
	if err := db.AutoMigrate(); err != nil {
		panic(err)
	}
	return db
}
