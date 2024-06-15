package handlers

import (
	go_redis "github.com/redis/go-redis/v9"

	"github.com/Danil-114195722/GoCurrencyCourseBot/redis"
)


var errorMessage string = "☠️ Возникла ошибка при выполнении. Попробуйте выйти в главное меню и попробовать ещё раз"
var	redisClient *go_redis.Client = redis.RedisClient()
