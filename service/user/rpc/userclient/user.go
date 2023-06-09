// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package userclient

import (
	"context"

	"XianShop/service/user/rpc/types/user"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CommonReply   = user.CommonReply
	LoginReq      = user.LoginReq
	RegisterReq   = user.RegisterReq
	SendCodeRep   = user.SendCodeRep
	SendCodeReply = user.SendCodeReply
	UserInfoReq   = user.UserInfoReq

	User interface {
		SendCode(ctx context.Context, in *SendCodeRep, opts ...grpc.CallOption) (*SendCodeReply, error)
		Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*CommonReply, error)
		Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*CommonReply, error)
		UserInfo(ctx context.Context, in *UserInfoReq, opts ...grpc.CallOption) (*CommonReply, error)
	}

	defaultUser struct {
		cli zrpc.Client
	}
)

func NewUser(cli zrpc.Client) User {
	return &defaultUser{
		cli: cli,
	}
}

func (m *defaultUser) SendCode(ctx context.Context, in *SendCodeRep, opts ...grpc.CallOption) (*SendCodeReply, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.SendCode(ctx, in, opts...)
}

func (m *defaultUser) Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*CommonReply, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.Register(ctx, in, opts...)
}

func (m *defaultUser) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*CommonReply, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.Login(ctx, in, opts...)
}

func (m *defaultUser) UserInfo(ctx context.Context, in *UserInfoReq, opts ...grpc.CallOption) (*CommonReply, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.UserInfo(ctx, in, opts...)
}
