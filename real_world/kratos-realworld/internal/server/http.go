package server

import (
	"context"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	v1 "kratos-realworld/api/realworld/v1"
	"kratos-realworld/internal/conf"
	"kratos-realworld/internal/pkg/middleware/auth"
	"kratos-realworld/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewWhitelistMatcher 白名单匹配器，用于跳过不需要中间件的接口
func NewWhitelistMatcher() selector.MatchFunc {
	whiteList := make(map[string]struct{})

	/*					以下摘自kratos文档：
	注意：定制中间件是通过 operation 匹配，并不是 http 本身的路由
	operation 是 HTTP 及 gRPC 统一的 gRPC path。
	gRPC path 的拼接规则为 /包名.服务名/方法名(/package.Service/Method)。
	*/
	whiteList["/realworld.v1.RealWorld/Register"] = struct{}{}
	whiteList["/realworld.v1.RealWorld/Login"] = struct{}{}
	return func(ctx context.Context, operation string) bool {
		if _, ok := whiteList[operation]; ok {
			return false
		}
		return true
	}
}

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, realWorldService *service.RealWorldService, jwt *conf.JWT, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			selector.Server(
				auth.JwtAuth(jwt.Secret),
			).Match(NewWhitelistMatcher()).Build(),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	v1.RegisterRealWorldHTTPServer(srv, realWorldService)
	return srv
}
