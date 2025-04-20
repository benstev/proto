package hagwclient

import (
	"context"

	"github.com/benstev/proto/hagw"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CallServiceReq  = hagw.CallServiceReq
	CallServiceResp = hagw.CallServiceResp
	FireEventReq    = hagw.FireEventReq
	FireEventResp   = hagw.FireEventResp
	HAState         = hagw.HAState
	StateReq        = hagw.StateReq
	StateResp       = hagw.StateResp
	SubscribeReq    = hagw.SubscribeReq
	SubscribeResp   = hagw.SubscribeResp

	Hagw interface {
		State(ctx context.Context, in *StateReq, opts ...grpc.CallOption) (*StateResp, error)
		Subscribe(ctx context.Context, in *SubscribeReq, opts ...grpc.CallOption) (hagw.Hagw_SubscribeClient, error)
		CallService(ctx context.Context, in *CallServiceReq, opts ...grpc.CallOption) (*CallServiceResp, error)
		FireEvent(ctx context.Context, in *FireEventReq, opts ...grpc.CallOption) (*FireEventResp, error)
	}

	defaultHagw struct {
		cli zrpc.Client
	}
)

func NewHagw(cli zrpc.Client) Hagw {
	return &defaultHagw{
		cli: cli,
	}
}

func (m *defaultHagw) State(ctx context.Context, in *StateReq, opts ...grpc.CallOption) (*StateResp, error) {
	client := hagw.NewHagwClient(m.cli.Conn())
	return client.State(ctx, in, opts...)
}

func (m *defaultHagw) Subscribe(ctx context.Context, in *SubscribeReq, opts ...grpc.CallOption) (hagw.Hagw_SubscribeClient, error) {
	client := hagw.NewHagwClient(m.cli.Conn())
	return client.Subscribe(ctx, in, opts...)
}

func (m *defaultHagw) CallService(ctx context.Context, in *CallServiceReq, opts ...grpc.CallOption) (*CallServiceResp, error) {
	client := hagw.NewHagwClient(m.cli.Conn())
	return client.CallService(ctx, in, opts...)
}

func (m *defaultHagw) FireEvent(ctx context.Context, in *FireEventReq, opts ...grpc.CallOption) (*FireEventResp, error) {
	client := hagw.NewHagwClient(m.cli.Conn())
	return client.FireEvent(ctx, in, opts...)
}
