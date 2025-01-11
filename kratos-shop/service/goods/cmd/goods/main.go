package main

import (
	"flag"
	"github.com/go-kratos/kratos/v2/registry"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	"os"
	"path/filepath"

	"goods/internal/conf"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	_ "go.uber.org/automaxprocs"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name string = "shop.goods.service"
	// Version is the version of the compiled software.
	Version string = "goods.v1"
	// flagconf is the config flag.
	flagconf string

	id, _ = os.Hostname()
)

func init() {
	flag.StringVar(&flagconf, "conf", "../../configs", "config path, eg: -conf config.yaml")
}

func newApp(logger log.Logger, gs *grpc.Server, consulRegistrar registry.Registrar) *kratos.App {
	return kratos.New(
		kratos.ID(id),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			gs,
		),
		kratos.Registrar(consulRegistrar),
	)
}

func main() {
	flag.Parse()
	logger := log.With(log.NewStdLogger(os.Stdout),
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"service.id", id,
		"service.name", Name,
		"service.version", Version,
		"trace.id", tracing.TraceID(),
		"span.id", tracing.SpanID(),
	)
	c := config.New(
		config.WithSource(
			file.NewSource(flagconf),
		),
	)
	defer c.Close()

	if err := c.Load(); err != nil {
		panic(err)
	}

	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}

	//引入Consul
	var rc conf.Registry
	if err := c.Scan(&rc); err != nil {
		panic(err)
	}

	// 配置jaeger链路追踪服务
	if err := setTraceProvider(bc.Trace.Endpoint); err != nil {
		panic(err)
	}

	//读取elasticsearch的cacert文件
	caPath, _ := filepath.Abs("./deploy/elasticsearch/configs/http_ca.crt")
	ca, err := os.ReadFile(caPath)
	if err != nil {
		panic(err)
	}
	bc.Data.ElasticSearch.CaCert = ca

	app, cleanup, err := wireApp(bc.Server, bc.Data, &rc, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}

func setTraceProvider(url string) error {
	// 设置 Jaeger 导出器，导出器负责将跟踪数据发送到 Jaeger 后端
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(url)))
	if err != nil {
		return err
	}

	//将 Jaeger 导出器设置为默认的跟踪后端
	tp := tracesdk.NewTracerProvider(
		// 设置采样率
		tracesdk.WithSampler(tracesdk.ParentBased(tracesdk.TraceIDRatioBased(1.0))),

		// 设置导出器
		tracesdk.WithBatcher(exp),

		// 设置资源信息
		tracesdk.WithResource(resource.NewSchemaless(
			semconv.ServiceNameKey.String(Name),
			attribute.String("env", "dev"),
		)),
	)
	otel.SetTracerProvider(tp)
	return nil
}
