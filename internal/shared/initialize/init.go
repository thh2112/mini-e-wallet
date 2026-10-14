package initialize

import (
	"context"
	"mini-e-wallet/internal/shared/global"
	"mini-e-wallet/pkg/config"
	"mini-e-wallet/pkg/logger"
)

func Init() {

	// Load Configuration
	cfg, err := config.LoadConfig("configs/config.yml")
	if err != nil {
		panic("Failed to load configuration: " + err.Error())
	}
	global.Config = cfg

	// Initialize Logger
	global.Logger = logger.GetLoggerInstance(cfg.Server.Mode)
	global.Logger.Info(context.Background(), "Logger initialized")

	// Initialize Database
	InitDatabase()

	// Initialize Redis
	InitRedis()

	// TODO: Initialize External API Clients
}
