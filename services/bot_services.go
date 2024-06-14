package services

import (
	telebot "gopkg.in/telebot.v3"

	"github.com/Danil-114195722/GoCurrencyCourseBot/settings"
)


// обработка неизвестной ошибки бота
func OnError(err error, context telebot.Context) {
	// лог с описанием ошибки
	settings.ErrorLog.Printf("Unknown Bot error handled: %v", err)
}
