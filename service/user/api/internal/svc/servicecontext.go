package svc

import (
	"XianShop/service/common"
	"XianShop/service/user/api/internal/config"
	"XianShop/service/user/model"
	"XianShop/service/user/rpc/userclient"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config    config.Config
	DbEngine  *gorm.DB
	UserModel model.UserModel
	BizRedis  *redis.Redis
	Rpc       userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	coon := sqlx.NewMysql(c.Mysql.DataSource)
	db := common.InitGorm(c.Mysql.DataSource)
	db.AutoMigrate(&model.User{})
	return &ServiceContext{
		Config:    c,
		UserModel: model.NewUserModel(coon, c.CacheRedis),
		BizRedis:  redis.New(c.BizRedis.Host, redis.WithPass(c.BizRedis.Pass)),
		Rpc:       userclient.NewUser(zrpc.MustNewClient(c.Rpc)),
		DbEngine:  db,
	}
}
