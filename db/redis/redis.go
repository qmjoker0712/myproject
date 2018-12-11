package redis

import (
	"myproject/types"

	"myproject/config"

	"github.com/s3dteam/go-toolkit/db/redis"
	"github.com/s3dteam/go-toolkit/log/logruslogger"
)

// DB redis service db
var DB *redis.RedisCacheService

//Init for redis
func Init(options *types.RedisOptions) {
	var redisService = &redis.RedisCacheService{}
	redisOptions := redis.RedisOptions{
		Host:        options.Host,
		Port:        options.Port,
		Password:    options.Password,
		IdleTimeout: options.IdleTimeout,
		MaxActive:   options.MaxActive,
		MaxIdle:     options.MaxIdle,
	}
	loger := logruslogger.GetLoggerWithOptions("ethaccessor.cache", &config.GetConfig().Log)
	redisService.Initialize(redisOptions, loger)

	DB = redisService
}
