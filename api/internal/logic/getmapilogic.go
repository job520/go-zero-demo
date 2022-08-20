package logic

import (
	"context"
	"errors"

	"demo/api/internal/svc"
	"demo/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMapiLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMapiLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMapiLogic {
	return &GetMapiLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMapiLogic) GetMapi(req *types.Req) (resp *types.Reply, err error) {
	id := req.Id
	detail, err := l.svcCtx.MtableModel.FindOne(l.ctx, id)
	if err != nil {
		return nil, errors.New("api: 未找到记录")
	}
	return &types.Reply{
		Id:   detail.Id,
		Name: detail.Name,
	}, nil
}
