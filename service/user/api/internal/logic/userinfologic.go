package logic

import (
	"XianShop/service/common/errorx"
	"XianShop/service/user/rpc/types/user"
	"context"
	"encoding/json"
	"fmt"

	"XianShop/service/user/api/internal/svc"
	"XianShop/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo(req *types.UserInfoReq) (resp *types.UserInfoResply, err error) {
	// todo: add your logic here and delete this line
	fmt.Println("开始api userinfo")
	userData, err := l.svcCtx.Rpc.UserInfo(l.ctx, &user.UserInfoReq{UserId: req.UserId})
	if err != nil {
		return nil, errorx.NewDefaultError(fmt.Sprintf("rpc查询userinfo错误：%v", err))
	}
	userInfoItem := &types.UserInfoItem{}
	userJsonErr := json.Unmarshal([]byte(userData.Data), &userInfoItem)
	if userJsonErr != nil {
		return nil, errorx.NewDefaultError("userinfo的json数据转换失败")
	}
	return &types.UserInfoResply{
		Code:    0,
		Message: "",
		Data:    userInfoItem,
	}, nil
}
