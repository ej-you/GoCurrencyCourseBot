package services

import (
	"strconv"

	telebot "gopkg.in/telebot.v3"

	"github.com/Danil-114195722/GoCurrencyCourseBot/settings"
)


// Обработка неизвестной ошибки бота
func OnError(err error, context telebot.Context) {
	// лог с описанием ошибки
	settings.ErrorLog.Printf("Unknown Bot error handled: %v", err)
}

// Возвращает строковый id юзера ТГ
func GetUserID(context telebot.Context) string {
	return strconv.FormatInt(context.Chat().ID, 10)
}
