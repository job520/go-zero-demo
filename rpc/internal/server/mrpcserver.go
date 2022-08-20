// Code generated by goctl. DO NOT EDIT!
// Source: mrpc.proto

package server

import (
	"context"

	"demo/rpc/internal/logic"
	"demo/rpc/internal/svc"
	"demo/rpc/types/mrpc"
)

type MrpcServer struct {
	svcCtx *svc.ServiceContext
	mrpc.UnimplementedMrpcServer
}

func NewMrpcServer(svcCtx *svc.ServiceContext) *MrpcServer {
	return &MrpcServer{
		svcCtx: svcCtx,
	}
}

func (s *MrpcServer) GetById(ctx context.Context, in *mrpc.Req) (*mrpc.Reply, error) {
	l := logic.NewGetByIdLogic(ctx, s.svcCtx)
	return l.GetById(in)
}
