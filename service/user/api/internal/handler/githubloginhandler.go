package handler

import (
	"fmt"
	"net/http"

	"XianShop/service/user/api/internal/logic"
	"XianShop/service/user/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GithubLoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewGithubLoginLogic(r.Context(), svcCtx)
		err := l.GithubLogin()
		authURL := fmt.Sprintf("https://github.com/login/oauth/authorize?client_id=%s&redirect_uri=%s", svcCtx.Config.Github.ClientID, svcCtx.Config.Github.RedirectUrl)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			// 重定向到回调
			http.Redirect(w, r, authURL, http.StatusTemporaryRedirect)
		}
	}
}
