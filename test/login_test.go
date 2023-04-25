package test

import (
	"XianShop/service/common"
	"testing"
)

func TestLogin(t *testing.T) {
	db := common.InitGorm("root:root@(127.0.0.1:3306)/xian-shop?charset=utf8mb4&parseTime=True&loc=Asia%2FShanghai")

}
