package global

import (
	"mini-e-wallet/pkg/config"
	"mini-e-wallet/pkg/logger"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	Logger logger.Logger
	Config *config.Config
	DB     *gorm.DB
	Redis  *redis.Client
)
