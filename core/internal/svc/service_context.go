package svc

import (
	"cloud-drive/core/internal/config"
	"cloud-drive/core/internal/middleware"
	"cloud-drive/core/models"
	"github.com/go-redis/redis/v9"
	"github.com/zeromicro/go-zero/rest"
	"xorm.io/xorm"
)

type ServiceContext struct {
	Config config.Config
	Engine *xorm.Engine
	RDB    *redis.Client
	Auth   rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Engine: models.Init(),
		RDB:    models.InitRedis(),
		Auth:   middleware.NewAuthMiddleware().Handle,
	}
}
