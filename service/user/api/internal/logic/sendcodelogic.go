package logic

import (
	"XianShop/service/user/api/internal/svc"
	"XianShop/service/user/api/internal/types"
	"XianShop/service/user/rpc/types/user"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendcodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendcodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendcodeLogic {
	return &SendcodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}

}

func (l *SendcodeLogic) Sendcode(req *types.RegisterByPhoneRep) (resp *types.RegisterByPhoneResply, err error) {
	// todo: add your logic here and delete this line
	cnt, cntErr := l.svcCtx.Rpc.SendCode(l.ctx, &user.SendCodeRep{
		UserPhone: req.UserPhone,
	})
	if cntErr != nil {
		return nil, cntErr
	}
	return &types.RegisterByPhoneResply{
		Code:    200,
		Message: "发送验证码success",
		VeCode:  cnt.VeCode,
	}, nil
}
