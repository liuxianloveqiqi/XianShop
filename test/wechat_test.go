package test

import (
	"fmt"
	"github.com/geiqin/thirdparty/oauth"
	"testing"
)

func TestWechat(t *testing.T) {
	wxConf := &oauth.AuthConfig{
		ClientId:     "your app_id",
		ClientSecret: "your app_secret",
		RedirectUrl:  "http://www.geiqin.com"}

	wxAuth := oauth.NewAuthQq(wxConf)
	fmt.Print(wxAuth.GetRedirectUrl("sate")) //获取第三方登录地址
	wxRes, err := wxAuth.GetToken("code")
	userInfo, _ := wxAuth.GetUserInfo(wxRes.AccessToken, wxRes.OpenId)

}
