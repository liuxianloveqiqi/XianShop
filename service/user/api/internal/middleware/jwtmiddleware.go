package middleware

import (
	"XianShop/service/utils"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
	"strings"
)

type JWTMiddleware struct {
}

func NewJWTMiddleware() *JWTMiddleware {
	return &JWTMiddleware{}
}

func (m *JWTMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	logx.Info("jwt middleware")
	return func(w http.ResponseWriter, r *http.Request) {
		// JWTAuthMiddleware implementation
		fmt.Println(000000)
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			w.WriteHeader(http.StatusBadRequest)
			err, _ := json.Marshal(errors.New("请求头中auth为空"))
			w.Write(err)
			fmt.Println("一错")
			return
		}
		fmt.Println(22222)
		parts := strings.Split(authHeader, " ")
		if !(len(parts) == 3 && parts[0] == "Bearer") {
			w.WriteHeader(http.StatusBadRequest)
			err, _ := json.Marshal(errors.New("请求头中auth格式有误"))
			w.Write(err)
			fmt.Println("二错")

			return
		}
		fmt.Println(333333)
		parseToken, isUpd, err := utils.ParseToken(parts[1], parts[2])
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			err, _ := json.Marshal(errors.New("无效的token"))
			w.Write(err)
			fmt.Println("三错")
			return
		}
		if isUpd {
			parts[1], parts[2] = utils.GetToken(parseToken.ID, parseToken.State)
			w.Header().Set("Authorization", fmt.Sprintf("Bearer %s %s", parts[1], parts[2]))
		}
		fmt.Println(44444)
		r = r.WithContext(context.WithValue(r.Context(), "userID", parseToken.ID))
		fmt.Println(parseToken.ID, "gggggg")
		// Passthrough to next handler
		next(w, r)
	}
}
