package logic

import (
	"context"

	"demo/rpc/internal/svc"
	"demo/rpc/types/test"

	"github.com/zeromicro/go-zero/core/logx"
)

type TestLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTestLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TestLogic {
	return &TestLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TestLogic) Test(in *test.Req) (*test.Resp, error) {
	logx.Info(in)
	return &test.Resp{
		Id:   1,
		Name: "lee",
	}, nil
}
