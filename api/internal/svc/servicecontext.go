package svc

import (
	"demo/api/internal/config"
	"demo/rpc/testclient"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	TestRpc testclient.Test
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		TestRpc: testclient.NewTest(zrpc.MustNewClient(c.TestRpc)),
	}
}
