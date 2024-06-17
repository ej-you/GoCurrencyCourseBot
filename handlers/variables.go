package handlers

import (
	go_redis "github.com/redis/go-redis/v9"
	telebot "gopkg.in/telebot.v3"

	"github.com/Danil-114195722/GoCurrencyCourseBot/redis"
	"github.com/Danil-114195722/GoCurrencyCourseBot/services"
)


var errorMessage string = "☠️ Возникла ошибка при выполнении. Попробуйте выйти в главное меню и попробовать ещё раз"
var notFoundCurrencyMessage string = "😬 Извините, по каким-то причинам курс валюты не найден. Попробуйте выйти в главное меню и попробовать ещё раз"
var invalidDateMessage string = "Неверная дата! Попробуйте ещё раз (формат: ДД/ММ/ГГГГ, с 01/01/1996 по текущий день)"
var	redisClient *go_redis.Client = redis.RedisClient()


// Возвращает true, если требуемое состояние юзера не равно действительному
func statusNotIs(context telebot.Context, needStatus string) bool {
	// получение состояния юзера
	status, err := redis.GetStatus(redisClient, services.GetUserID(context))
	if err != nil {
		return true
	}

	if needStatus == status {
		return false
	} else {
		return true
	}
}
