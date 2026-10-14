package initialize

import (
	"context"
	"mini-e-wallet/internal/shared/global"
)

func closeDB(ctx context.Context) {
	global.Logger.Info(ctx, "Closing database")

	if db, ok := global.DB.DB(); ok == nil {
		if err := db.Close(); err != nil {
			global.Logger.Error(ctx, "Error closing database connection: "+err.Error())
		} else {
			global.Logger.Info(ctx, "Database connection closed")
		}
	}
}

func closeRedis(ctx context.Context) {
	global.Logger.Info(ctx, "Closing redis")

	if global.Redis != nil {
		if err := global.Redis.Close(); err != nil {
			global.Logger.Error(ctx, "Error closing redis connection: "+err.Error())
		} else {
			global.Logger.Info(ctx, "Redis connection closed")
		}
	}
}

func Close(ctx context.Context) {
	closeDB(ctx)
	closeRedis(ctx)
}
