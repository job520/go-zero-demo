package logic

import (
	"context"
	"errors"

	"demo/rpc/internal/svc"
	"demo/rpc/types/mrpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetByIdLogic {
	return &GetByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetByIdLogic) GetById(in *mrpc.Req) (*mrpc.Reply, error) {
	id := in.Id
	detail, err := l.svcCtx.MtableModel.FindOne(l.ctx, id)
	if err != nil {
		return nil, errors.New("rpc-服务端: 未找到记录")
	}
	return &mrpc.Reply{
		Id:   detail.Id,
		Name: detail.Name,
	}, nil
}
