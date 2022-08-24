// Code generated by goctl. DO NOT EDIT!
// Source: test.proto

package testclient

import (
	"context"

	"demo/rpc/types/test"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	Req  = test.Req
	Resp = test.Resp

	Test interface {
		Test(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Resp, error)
	}

	defaultTest struct {
		cli zrpc.Client
	}
)

func NewTest(cli zrpc.Client) Test {
	return &defaultTest{
		cli: cli,
	}
}

func (m *defaultTest) Test(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Resp, error) {
	client := test.NewTestClient(m.cli.Conn())
	return client.Test(ctx, in, opts...)
}
