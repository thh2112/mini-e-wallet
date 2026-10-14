package initialize

import (
	"context"
	"mini-e-wallet/internal/shared/global"
	"mini-e-wallet/pkg/config"

	"log"
	"os"
	"sync"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	dbOnce sync.Once
	db     *gorm.DB
)

func setup(config *config.DatabaseConfig) *gorm.DB {
	dbOnce.Do(func() {
		ctx := context.Background()
		newLogger := logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold: time.Second,
				LogLevel:      logger.Info,
				Colorful:      true,
			},
		)

		gormDB, err := gorm.Open(postgres.Open(config.Url), &gorm.Config{
			Logger: newLogger,
		})

		if err != nil {
			log.Fatal("Failed to connect to database:", err)
		}

		sqlDB, err := gormDB.DB()
		if err != nil {
			panic("Failed to get sqlDB: " + err.Error())
		}

		sqlDB.SetMaxIdleConns(config.MaxIdleConns)
		sqlDB.SetMaxOpenConns(config.MaxOpenConns)
		sqlDB.SetConnMaxLifetime(config.ConnMaxLifetime)

		db = gormDB
		global.Logger.Info(ctx, "Database initialized")

	})

	return db
}

func InitDatabase() {
	global.DB = setup(&global.Config.Database)
	global.Logger.Info(context.Background(), "Database initialized")
}
