package logic

import (
	"context"
	"demo/rpc/mrpcclient"
	"errors"

	"demo/api/internal/svc"
	"demo/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMrpcLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMrpcLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMrpcLogic {
	return &GetMrpcLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMrpcLogic) GetMrpc(req *types.Req) (resp *types.Reply, err error) {
	id := req.Id
	detail, err := l.svcCtx.Mrpc.GetById(l.ctx, &mrpcclient.Req{
		Id: id,
	})
	if err != nil {
		return nil, errors.New("rpc-客户端: 未找到记录")
	}
	return &types.Reply{
		Id:   detail.Id,
		Name: detail.Name,
	}, nil
}
