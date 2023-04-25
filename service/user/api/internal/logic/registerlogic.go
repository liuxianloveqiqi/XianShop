package logic

import (
	"XianShop/service/user/rpc/types/user"
	"context"

	"XianShop/service/user/api/internal/svc"
	"XianShop/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.CommonResply, err error) {
	// todo: add your logic here and delete this line
	cnt, cntErr := l.svcCtx.Rpc.Register(l.ctx, &user.RegisterReq{
		UserPhone: req.UserPhone,
		VeCode:    req.VeCode,
	})
	if cntErr != nil {
		return nil, cntErr
	}

	return &types.CommonResply{
		Code:    cnt.Code,
		Message: cnt.Message,
		Data:    cnt.Data,
	}, nil
}
