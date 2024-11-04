package main

import (
	"time"

	telebot "gopkg.in/telebot.v3"

	"github.com/ej-you/GoCurrencyCourseBot/middlewares"
	"github.com/ej-you/GoCurrencyCourseBot/services"
	"github.com/ej-you/GoCurrencyCourseBot/settings"

	"github.com/ej-you/GoCurrencyCourseBot/handlers"
	"github.com/ej-you/GoCurrencyCourseBot/keyboards"
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

	// создание группы хэндлеров и добавление middleware для лога каждой введённой команды (нажатой кнопки, аналогичной команде)
	commandsHandlers := bot.Group()
	commandsHandlers.Use(middlewares.CommandsLogger)

	// инициализация хендлеров
	commandsHandlers.Handle("/start", handlers.StartHandler)

	commandsHandlers.Handle("/home", handlers.HomeHandler)
	commandsHandlers.Handle("/cancel", handlers.HomeHandler)
	commandsHandlers.Handle(&keyboards.BtnBackToHome, handlers.HomeHandler)
	commandsHandlers.Handle(&keyboards.BtnCourseBackToHome, handlers.HomeHandler)

	commandsHandlers.Handle("/help", handlers.HelpHandler)
	
	commandsHandlers.Handle("/currencies", handlers.CurrenciesHandler)
	commandsHandlers.Handle(&keyboards.BtnCurrrencies, handlers.CurrenciesHandler)
	
	commandsHandlers.Handle("/course", handlers.CourseHandler)
	commandsHandlers.Handle(&keyboards.BtnCurrrencyCourse, handlers.CourseHandler)
	commandsHandlers.Handle(&keyboards.BtnGetCourseAgain, handlers.CourseHandler)
	
	bot.Handle(&keyboards.BtnActualCourse, handlers.ActualCourseHandler)
	bot.Handle(&keyboards.BtnHistoricalCourse, handlers.HistoricalCourseHandler)

	bot.Handle(telebot.OnText, handlers.CourseDialogHandler)
	bot.Handle(telebot.OnCallback, handlers.CourseDialogHandler)

	// запуск бота
	settings.InfoLog.Printf("Start bot %s...", bot.Me.Username)
	bot.Start()
}
