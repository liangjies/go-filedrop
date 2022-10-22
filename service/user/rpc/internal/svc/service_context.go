package svc

import (
	"go-filedrop/service/user/rpc/internal/config"
	"go-filedrop/service/user/rpc/models"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
	RDB    *redis.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		DB:     models.InitDB(c.MySQL.DataSource),
		RDB:    models.InitRedis(c),
	}
}
