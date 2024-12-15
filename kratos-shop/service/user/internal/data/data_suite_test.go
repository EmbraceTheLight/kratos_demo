package data_test

import (
	"context"
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
	"github.com/pkg/errors"
	"user/internal/conf"

	"gorm.io/gorm"
	"testing"
	"user/internal/data"
)

func TestData(t *testing.T) {
	// Ginkgo 测试通过调用Fail函数来表示失败
	//使用RegisterFailHandler加你个此函数传递给Gomega。这是Ginkgo和gomega之间的唯一连接点。
	gomega.RegisterFailHandler(ginkgo.Fail)
	//通知Ginkgo启动测试套件。如果任何specs失败，Ginkgo将自动生成testing.T失败
	ginkgo.RunSpecs(t, "test biz data")
}

var (
	cleaner func()          //定义删除mysql容器的回调函数
	DB      *data.Data      //用于测试的data
	ctx     context.Context //用于测试的context
)

// 初始化db。自动建表
func initialize(db *gorm.DB) error {
	err := db.AutoMigrate(
		&data.User{},
	)
	return errors.WithStack(err)
}

// ginkgo使用 BeforeSuite 来为spec设置状态
var _ = ginkgo.BeforeSuite(func() {
	//连接docker容器船舰的mysql
	conn, f := data.DockerMySQL("mysql", "latest")
	cleaner = f
	config := &conf.Data{
		Mysql: &conf.Data_Mysql{
			Driver: "mysql",
			Source: conn,
		},
	}
	db := data.NewMySQL(config)
	mysqlDB, _, err := data.NewData(config, db, nil, nil)
	if err != nil {
		return
	}
	DB = mysqlDB
	err = initialize(db)
	if err != nil {
		return
	}
	gomega.Expect(err).NotTo(gomega.HaveOccurred())
})

// 测试结束后，通过回调函数，关闭并删除docker创建的容器
var _ = ginkgo.AfterSuite(func() {
	cleaner()
})
