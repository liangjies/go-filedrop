package svc

import (
	"go-filedrop/service/store/api/internal/config"
	"go-filedrop/service/store/api/internal/middleware"
	"go-filedrop/service/store/api/models"
	"go-filedrop/service/user/rpc/userclient"

	"github.com/go-redis/redis/v8"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config  config.Config
	DB      *gorm.DB
	RDB     *redis.Client
	Auth    rest.Middleware
	UserRpc userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		DB:      models.InitDB(c.MySQL.DataSource),
		RDB:     models.InitRedis(c),
		Auth:    middleware.NewAuthMiddleware().Handle,
		UserRpc: userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
