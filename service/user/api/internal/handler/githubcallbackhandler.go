package handler

import (
	"XianShop/service/auth"
	"XianShop/service/auth/auth_model"
	"context"
	"fmt"
	"net/http"

	"XianShop/service/user/api/internal/logic"
	"XianShop/service/user/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GithubCallbackHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 从查询参数中获取授权码
		code := r.URL.Query().Get("code")
		// 交换授权码获取访问令牌
		tokenAuthUrl := fmt.Sprintf(
			"https://github.com/login/oauth/access_token?client_id=%s&client_secret=%s&code=%s",
			svcCtx.Config.Github.ClientID, svcCtx.Config.Github.ClientSecret, code)
		// 获取 token
		var token *auth_model.GithubToken
		var err error
		if token, err = auth.GetGithubToken(tokenAuthUrl); err != nil {
			fmt.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Printf("------%v", token)
		// 通过token，获取用户信息
		user0, err := auth.GetUserInfo(token)
		if err != nil {
			fmt.Println("获取用户信息失败，错误信息为:", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		r = r.WithContext(context.WithValue(r.Context(), "github_id", user0.GithubId))
		r = r.WithContext(context.WithValue(r.Context(), "github_nickname", user0.Nickname))
		l := logic.NewGithubCallbackLogic(r.Context(), svcCtx)
		resp, err := l.GithubCallback()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
