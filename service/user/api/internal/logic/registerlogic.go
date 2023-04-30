package logic

import (
	"XianShop/service/common/errorx"
	"XianShop/service/user/rpc/types/user"
	"XianShop/service/utils"
	"context"
	"fmt"
	"github.com/google/uuid"

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

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.TokenResply, err error) {
	// todo: add your logic here and delete this line
	cnt, cntErr := l.svcCtx.Rpc.Register(l.ctx, &user.RegisterReq{
		UserPhone: req.UserPhone,
		VeCode:    req.VeCode,
	})
	if cntErr != nil {
		return nil, cntErr
	}
	accessTokenString, refreshTokenString := utils.GetToken(cnt.UserId, uuid.New().String())
	if accessTokenString == "" || refreshTokenString == "" {
		return nil, errorx.NewDefaultError("jwt生成错误")
	}
	fmt.Println(cnt, 4234324324)
	return &types.TokenResply{
		Code:         cnt.Code,
		Message:      cnt.Message,
		UserId:       cnt.UserId,
		AccessToken:  accessTokenString,
		RefreshToken: refreshTokenString,
	}, nil
}
