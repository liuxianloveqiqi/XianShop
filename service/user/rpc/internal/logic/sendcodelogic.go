package logic

import (
	"XianShop/service/utils"
	"context"
	"time"

	"XianShop/service/user/rpc/internal/svc"
	"XianShop/service/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendCodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendCodeLogic {
	return &SendCodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SendCodeLogic) SendCode(in *user.SendCodeRep) (*user.SendCodeReply, error) {
	// todo: add your logic here and delete this line

	err := utils.DefaultGetValidParams(l.ctx, in)
	if err != nil {
		return nil, err
	}
	l.svcCtx.Rdb.Set(l.ctx, "8888", 989890, 3*time.Minute)
	vecode := utils.SMS(in.UserPhone, l.svcCtx.Config.Credential.SecretId, l.svcCtx.Config.Credential.SecretKey, l.ctx, l.svcCtx.Rdb)
	return &user.SendCodeReply{
		VeCode: vecode,
	}, nil
}
