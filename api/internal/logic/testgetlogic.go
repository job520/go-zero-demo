package logic

import (
	"context"
	"demo/rpc/testclient"

	"demo/api/internal/svc"
	"demo/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TestGetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTestGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TestGetLogic {
	return &TestGetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TestGetLogic) TestGet(req *types.Req) (resp *types.Resp, err error) {
	// 调用 go-zero-rpc
	ret, err := l.svcCtx.TestRpc.Test(l.ctx, &testclient.Req{})
	logx.Info(ret, err)
	return &types.Resp{
		Id:   "1",
		Name: l.svcCtx.Config.TestValue,
	}, nil
}
