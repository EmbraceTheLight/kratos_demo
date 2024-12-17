package data

import (
	"context"
	"github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	consulapi "github.com/hashicorp/consul/api"
	userv1 "shop/api/service/user/v1"
	"shop/internal/conf"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,
	NewUserRepo,
	NewUserServiceClient,
	NewRegistrar,
	NewDiscovery)

// Data .
type Data struct {
	log *log.Helper
	uc  userv1.UserClient
}

// NewData .
func NewData(
	c *conf.Data,
	uc userv1.UserClient,
	logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{
		log: log.NewHelper(log.With(logger, "module", "data")),
		uc:  uc,
	}, cleanup, nil
}

// NewUserServiceClient 链接用户服务
func NewUserServiceClient(
	ac *conf.Auth,
	sr *conf.Service,
	rr registry.Discovery) userv1.UserClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint(sr.User.Endpoint),
		grpc.WithDiscovery(rr),
		grpc.WithMiddleware(
			recovery.Recovery(),
		),
		grpc.WithTimeout(2*time.Second),
	)
	if err != nil {
		panic(err)
	}
	c := userv1.NewUserClient(conn)
	return c
}

// NewRegistrar 添加consul注册中心
func NewRegistrar(conf *conf.Registry) registry.Registrar {
	c := consulapi.DefaultConfig()
	c.Address = conf.Consul.Address
	c.Scheme = conf.Consul.Scheme
	cli, err := consulapi.NewClient(c)
	if err != nil {
		panic(err)
	}
	r := consul.New(cli, consul.WithHealthCheck(true))
	return r
}

// NewDiscovery 服务发现
func NewDiscovery(conf *conf.Registry) registry.Discovery {
	c := consulapi.DefaultConfig()
	c.Address = conf.Consul.Address
	c.Scheme = conf.Consul.Scheme
	cli, err := consulapi.NewClient(c)
	if err != nil {
		panic(err)
	}
	r := consul.New(cli, consul.WithHealthCheck(true))
	return r
}
