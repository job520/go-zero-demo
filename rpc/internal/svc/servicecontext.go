package svc

import (
	"demo/model"
	"demo/rpc/internal/config"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config      config.Config
	MtableModel model.MtableModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:      c,
		MtableModel: model.NewMtableModel(conn, c.CacheRedis),
	}
}
