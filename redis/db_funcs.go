package redis

import (
	"context"
	// "fmt"

	go_redis "github.com/redis/go-redis/v9"

	"github.com/Danil-114195722/GoCurrencyCourseBot/settings"
)


func SetStatus(redisClient *go_redis.Client, userId, status string) error {
	// пустой контекст для выполнения запросов к redis
	ctx := context.Background()

	// установка ключа-значения в redis
	err := redisClient.Set(ctx, userId, status, 0).Err()
	if err != nil {
	    settings.ErrorLog.Println("Failed to set status in Redis:", err)
	    return err
	}

	// лог о смене состояния юзера
	settings.InfoLog.Printf("Set %q status of user %s in redis", status, userId)
	return nil
}

func GetStatus(redisClient *go_redis.Client, userId string) (string, error) {
	var status string

	// пустой контекст для выполнения запросов к redis
	ctx := context.Background()

	// получение значения из redis по ключу
	status, err := redisClient.Get(ctx, userId).Result()
	if err != nil {
	    settings.ErrorLog.Println("Failed to get status from Redis:", err)
	    return status, err
	}
	return status, nil
}
