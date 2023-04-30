package logic

import (
	"XianShop/service/user/model"
	"XianShop/service/utils"
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"time"

	"XianShop/service/user/api/internal/svc"
	"XianShop/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GithubCallbackLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGithubCallbackLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GithubCallbackLogic {
	return &GithubCallbackLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GithubCallbackLogic) GithubCallback() (resp *types.TokenResply, err error) {
	users, err := l.svcCtx.UserModel.FindUserBy(l.svcCtx.DbEngine, "github_id", l.ctx.Value("github_id"))
	if err != nil {
		return nil, err
	}

	if len(users) == 0 {
		fmt.Println("该用户为githubu新用户，开始注册")
		// 新建用户
		user1 := model.User{
			PassWord:   utils.Md5Password(utils.GeneratePassword(10), "liuxian"),
			UserNick:   l.ctx.Value("github_nickname").(string),
			GithubId:   l.ctx.Value("github_id").(string),
			UserSex:    2,
			CreateTime: time.Now(),
			UpdateTime: time.Now(),
		}
		l.svcCtx.DbEngine.Create(&user1)
		accessTokenString, refreshTokenString := utils.GetToken(user1.UserId, uuid.New().String())
		if accessTokenString == "" || refreshTokenString == "" {
			return nil, errors.New("生成jwt错误")
		}
		return &types.TokenResply{
			Code:         200,
			Message:      "success!",
			UserId:       user1.UserId,
			AccessToken:  accessTokenString,
			RefreshToken: refreshTokenString,
		}, nil
	}
	accessTokenString, refreshTokenString := utils.GetToken(users[0].UserId, uuid.New().String())
	if accessTokenString == "" || refreshTokenString == "" {
		return nil, errors.New("生成jwt错误")
	}
	return &types.TokenResply{
		Code:         200,
		Message:      "success!",
		UserId:       users[0].UserId,
		AccessToken:  accessTokenString,
		RefreshToken: refreshTokenString,
	}, nil
}
