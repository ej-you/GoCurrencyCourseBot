package main

import (
	"time"

	telebot "gopkg.in/telebot.v3"

	"github.com/Danil-114195722/GoCurrencyCourseBot/services"
	"github.com/Danil-114195722/GoCurrencyCourseBot/settings"

	"github.com/Danil-114195722/GoCurrencyCourseBot/handlers"
	"github.com/Danil-114195722/GoCurrencyCourseBot/keyboards"
)


func main() {
	// настройки бота
	pref := telebot.Settings{
		Token:  settings.BotToken,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
		// Verbose: true,
		OnError: services.OnError,
	}

	// инициализация бота
	bot, err := telebot.NewBot(pref)
	settings.DieIf(err)

	// инициализация клавиатур
	keyboards.InitKeyboards()

	// инициализация хендлеров
	bot.Handle("/start", handlers.StartHandler)

	bot.Handle("/home", handlers.HomeHandler)
	bot.Handle("/cancel", handlers.HomeHandler)
	bot.Handle(&keyboards.BtnBackToHome, handlers.HomeHandler)

	bot.Handle("/help", handlers.HelpHandler)
	
	bot.Handle("/course", handlers.CourseHandler)
	bot.Handle(&keyboards.BtnCurrrencyCourse, handlers.CourseHandler)
	
	bot.Handle("/currencies", handlers.CurrenciesHandler)
	bot.Handle(&keyboards.BtnCurrrencies, handlers.CurrenciesHandler)

	// запуск бота
	settings.InfoLog.Printf("Start bot %s...", bot.Me.Username)
	bot.Start()
}
