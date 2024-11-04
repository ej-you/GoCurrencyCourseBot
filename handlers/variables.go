package handlers

import (
	"fmt"

	go_redis "github.com/redis/go-redis/v9"
	telebot "gopkg.in/telebot.v3"

	"github.com/ej-you/GoCurrencyCourseBot/redis"
	"github.com/ej-you/GoCurrencyCourseBot/services"
	"github.com/ej-you/GoCurrencyCourseBot/settings"
)


var errorMessage string = "☠️ Возникла ошибка при выполнении. Попробуйте выйти в главное меню и попробовать ещё раз"
var notFoundCurrencyMessage string = "😬 Извините, по каким-то причинам курс валюты не найден. Попробуйте выйти в главное меню и попробовать ещё раз"
var invalidDateMessage string = fmt.Sprintf("Неверная дата! Попробуйте ещё раз (формат: ДД/ММ/ГГГГ, с %s по текущий день)", settings.LowDate)
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
