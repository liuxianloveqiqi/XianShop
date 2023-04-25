package common

import (
	"github.com/zeromicro/go-zero/core/logx"
	"testing"
)

func TestInitGorm(t *testing.T) {
	InitGorm("root:root@(127.0.0.1:3306)/xian-shop?charset=utf8mb4&parseTime=True&loc=Asia%2FShanghai")
	logx.Info("hello world")

	logx.Info("hello wadld")
}
