package logic

import (
	"XianShop/service/common/errorx"
	"XianShop/service/user/rpc/internal/svc"
	"XianShop/service/user/rpc/types/user"
	"context"
	"encoding/json"
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserInfoLogic) UserInfo(in *user.UserInfoReq) (*user.CommonReply, error) {
	// todo: add your logic here and delete this line
	fmt.Println("rpc userinfo")
	user0, err := l.svcCtx.UserModel.FindOne(l.ctx, in.UserId)
	if err != nil {
		return nil, errorx.NewDefaultError("找不到用户")
	}
	userJson, err := json.Marshal(user0)
	if err != nil {
		return nil, errorx.NewDefaultError("数据转换失败")
	}
	return &user.CommonReply{
		Code:    200,
		Message: "或许用户信息成功！",
		Data:    string(userJson),
	}, nil
}
