package redis

import (
	"fmt"
	"context"

	go_redis "github.com/redis/go-redis/v9"

	"github.com/ej-you/GoCurrencyCourseBot/settings"
)


// установка состояния юзера
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
	// settings.InfoLog.Printf("Set %q status of user %s in redis", status, userId)
	return nil
}

// получение состояния юзера
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

// установка выбранной валюты юзером
func SetChosenCurrency(redisClient *go_redis.Client, userId, chosenCurrency string) error {
	// пустой контекст для выполнения запросов к redis
	ctx := context.Background()

	// создание уникального ключа для записи в redis
	uniqueKey := fmt.Sprintf("%s_currency", userId)

	// установка ключа-значения в redis
	err := redisClient.Set(ctx, uniqueKey, chosenCurrency, 0).Err()
	if err != nil {
	    settings.ErrorLog.Println("Failed to set chosenCurrency in Redis:", err)
	    return err
	}

	// лог о выборе валюты юзером
	settings.InfoLog.Printf("Set %q chosenCurrency of user %s in redis", chosenCurrency, userId)
	return nil
}

// получение выбранной валюты юзером
func GetChosenCurrency(redisClient *go_redis.Client, userId string) (string, error) {
	var chosenCurrency string

	// пустой контекст для выполнения запросов к redis
	ctx := context.Background()

	// воссоздание уникального ключа для получениия валюты из redis
	uniqueKey := fmt.Sprintf("%s_currency", userId)

	// получение значения из redis по ключу
	chosenCurrency, err := redisClient.Get(ctx, uniqueKey).Result()
	if err != nil {
	    settings.ErrorLog.Println("Failed to get chosenCurrency from Redis:", err)
	    return chosenCurrency, err
	}
	return chosenCurrency, nil
}
