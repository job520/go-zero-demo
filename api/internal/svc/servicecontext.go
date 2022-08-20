package svc

import (
	"demo/api/internal/config"
	"demo/model"
	"demo/rpc/mrpcclient"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config      config.Config
	MtableModel model.MtableModel
	Mrpc        mrpcclient.Mrpc
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:      c,
		MtableModel: model.NewMtableModel(conn, c.CacheRedis),
		Mrpc:        mrpcclient.NewMrpc(zrpc.MustNewClient(c.Mrpc)),
	}
}
