package main

import (
	"flag"
	"fmt"
	"go-zero-demo/mall/user/rpc/internal/config"
	"go-zero-demo/mall/user/rpc/internal/server"
	"go-zero-demo/mall/user/rpc/internal/svc"
	"go-zero-demo/mall/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/user.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	svr := server.NewUserServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		user.RegisterUserServer(grpcServer, svr)

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	// 注册到nacos上 注册服务
	/*sc := []constant.ServerConfig{
		*constant.NewServerConfig("127.0.0.1", 8848),
	}
	cc := &constant.ClientConfig{
		NamespaceId:         "public",
		TimeoutMs:           50000,
		NotLoadCacheAtStart: true,
		LogLevel:            "error",
	}

	opts := nacos.NewNacosConfig("user.rpc", c.ListenOn, sc, cc)
	_ = nacos.RegisterService(opts)*/

	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
