package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	Mysql struct {
		DataSource string
	}
	CacheRedis cache.CacheConf
	Rpc        zrpc.RpcClientConf
	Github     struct {
		ClientID     string
		RedirectUrl  string
		ClientSecret string
	}
}
