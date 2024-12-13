package data

import (
	"database/sql"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/ory/dockertest/v3"
	"time"
)

func DockerMySQL(img, version string) (string, func()) {
	return innerDockerMySQL(img, version)
}
func innerDockerMySQL(img, version string) (string, func()) {
	// uses a sensible default on windows (tcp/http) and linux/osx (socket)
	// 创建一个docker连接池，用于拉取、删除镜像
	pool, err := dockertest.NewPool("")
	pool.MaxWait = 2 * time.Minute
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	//拉取img镜像，并基于该镜像创建并启动一个容器
	resource, err := pool.Run(img, version,
		[]string{
			"MYSQL_ROOT_PASSWORD=secret",
			"MYSQL_ROOT_HOST=%",
		})
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	conStr := fmt.Sprintf("root:secret@(localhost:%s)/mysql?parseTime=true", resource.GetPort("3306/tcp"))
	// pool.Retry使用指数退避算法来重试操作，直到成功或达到最大重试次数
	// 这里操作为打开数据库连接
	if err := pool.Retry(func() error {
		var err error
		db, err := sql.Open("mysql", conStr)
		if err != nil {
			return err
		}
		return db.Ping()
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	// 回调函数用于关闭并删除容器
	return conStr, func() {
		if err = pool.Purge(resource); err != nil {
			log.Fatalf("Could not purge resource: %s", err)
		}
	}
}
