package redis


import (
	"context"

	go_redis "github.com/redis/go-redis/v9"

	"github.com/ej-you/GoCurrencyCourseBot/settings"
)


// получение клиента redis
func RedisClient() *go_redis.Client {
	settings.InfoLog.Printf("Connect to redis on %s...", settings.RedisAddr)

	// создание нового клиента
	redisClient := go_redis.NewClient(&go_redis.Options{
		Addr: settings.RedisAddr,
		Password: settings.RedisPassword,
		DB: 0,
	})

	// пустой контекст для выполнения запросов к redis
	ctx := context.Background()

	// проверка подключения
	pong, err := redisClient.Ping(ctx).Result()
	settings.DieIf(err)

	settings.InfoLog.Printf("Successfully connected to redis: PING - %s", pong)

	return redisClient
}
