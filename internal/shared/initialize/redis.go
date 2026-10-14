package initialize

import (
	"context"
	"mini-e-wallet/internal/shared/global"
	"sync"

	"github.com/redis/go-redis/v9"
)

var (
	redisOnce sync.Once
)

func InitRedis() {
	redisOnce.Do(func() {
		ctx := context.Background()

		opt, err := redis.ParseURL(global.Config.Redis.Url)
		if err != nil {
			global.Logger.Error(ctx, "Failed to parse Redis URL: "+err.Error())
			return
		}

		client := redis.NewClient(opt)

		if err := client.Ping(ctx).Err(); err != nil {
			global.Logger.Error(ctx, "Failed to connect to Redis: "+err.Error())
			return
		}

		global.Redis = client
		global.Logger.Info(ctx, "Redis connected")
	})
}
