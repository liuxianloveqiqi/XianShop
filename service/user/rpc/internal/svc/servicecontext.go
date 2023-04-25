package svc

import (
	"XianShop/service/common"
	"XianShop/service/user/model"
	"XianShop/service/user/rpc/internal/config"
	"github.com/redis/go-redis/v9"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config    config.Config
	UserModel model.UserModel
	Rdb       *redis.Client
	DbEngine  *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	coon := sqlx.NewMysql(c.Mysql.DataSource)
	db := common.InitGorm(c.Mysql.DataSource)
	rdb := common.InitRedis(c.Redis.Host, c.Redis.Pass, c.Redis.DB)
	db.AutoMigrate(&model.User{})
	return &ServiceContext{
		Config:    c,
		UserModel: model.NewUserModel(coon, c.CacheRedis),
		Rdb:       rdb,
		DbEngine:  db,
	}
}
