package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"go-zero-cloud/mall/user/rpc/user"
	"go-zero-cloud/order/api/internal/config"
)

type ServiceContext struct {
	Config  config.Config
	UserRpc user.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserRpc: user.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
