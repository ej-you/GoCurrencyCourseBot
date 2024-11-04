package handlers

import (
	"fmt"
	"strings"

	telebot "gopkg.in/telebot.v3"

	"github.com/ej-you/GoCurrencyCourseBot/currency_api"
	"github.com/ej-you/GoCurrencyCourseBot/keyboards"
	"github.com/ej-you/GoCurrencyCourseBot/redis"
	"github.com/ej-you/GoCurrencyCourseBot/services"
	"github.com/ej-you/GoCurrencyCourseBot/settings"
)


// команда /course
func CourseHandler(context telebot.Context) error {
	if statusNotIs(context, "home") && statusNotIs(context, "chosen_action") && statusNotIs(context, "gotten_historical_course") {
		return nil
	}

	newStatus := "started"
	// установка состояния юзера
	err := redis.SetStatus(redisClient, services.GetUserID(context), newStatus)
	if err != nil {
		return context.Send(errorMessage, keyboards.BackToHomeInlineKeyboard)
	}

	msgText := `Выберите валюту, курс которой хотите узнать 👇

Если вам не понятны коды, то их расшифровки можно посмотреть в разделе "список валют" в главном меню.
Или просто введите /currencies`

	return context.Send(msgText, keyboards.CurrenciesListInlineKeyboard)
}

// диалог после команды /course, распределитель обработчиков (с неявным сигналом на обработку) по состояниям
func CourseDialogHandler(context telebot.Context) error {
	// получение состояния юзера
	status, err := redis.GetStatus(redisClient, services.GetUserID(context))
	if err != nil {
		return context.Send(errorMessage, keyboards.BackToHomeInlineKeyboard)
	}

	switch status {
		case "home":
			return nil
		// после выбора вылюты  обрабатываем её и предлагаем действие
		case "started":
			return ChooseCurrency(context)
		case "chosen_currency":
			return nil
		case "chosen_action":
			return nil
		// обрабатываем ввод даты для получения курса валюты за неё
		case "choose_date":
			return ChooseDate(context)
		case "gotten_historical_course":
			return nil
		default:
			return nil
	}
}


// обработка выбора юзером валюты
func ChooseCurrency(context telebot.Context) error {
	if statusNotIs(context, "started") {
		return nil
	}

	newStatus := "chosen_currency"
	// установка состояния юзера
	err := redis.SetStatus(redisClient, services.GetUserID(context), newStatus)
	if err != nil {
		return context.Send(errorMessage, keyboards.BackToHomeInlineKeyboard)
	}

	// достаём выбранную валюту из callback'а (сразу очищаем от префикса \f)
	chosenCurrency := strings.TrimPrefix(context.Callback().Data, "\f")

	// запись выбранной юзером валюты в redis
	err = redis.SetChosenCurrency(redisClient, services.GetUserID(context), chosenCurrency)
	if err != nil {
		return context.Send(errorMessage, keyboards.BackToHomeInlineKeyboard)
	}

	msgText := fmt.Sprintf("Выбранная валюта: %s\n\nВыберите действие 👇", chosenCurrency)
	return context.Send(msgText, keyboards.ActionInlineKeyboard)
}

// обработка выдачи актуального курса валюты юзеру
func ActualCourseHandler(context telebot.Context) error {
	if statusNotIs(context, "chosen_currency") {
		return nil
	}

	newStatus := "chosen_action"
	// установка состояния юзера
	err := redis.SetStatus(redisClient, services.GetUserID(context), newStatus)
	if err != nil {
		return context.Send(errorMessage, keyboards.BackToHomeInlineKeyboard)
	}

	// получение выбранной юзером валюты из redis
	chosenCurrencyCode, err := redis.GetChosenCurrency(redisClient, services.GetUserID(context))
	if err != nil {
		return context.Send(errorMessage, keyboards.BackToHomeInlineKeyboard)
	}

	// получение актуального курса из API
	latestCurse, err := currency_api.GetLatestCourse(chosenCurrencyCode)
	if err != nil {
		return context.Send(notFoundCurrencyMessage, keyboards.BackToHomeInlineKeyboard)
	}

	msgText := fmt.Sprintf("Актуальный курс %s - %.2f ₽\n\nВыберите действие 👇", chosenCurrencyCode, latestCurse)
	return context.Send(msgText, keyboards.GottenCourseInlineKeyboard)
}

// запрос даты для получения курса валюты за неё
func HistoricalCourseHandler(context telebot.Context) error {
	if statusNotIs(context, "chosen_currency") {
		return nil
	}

	newStatus := "choose_date"
	// установка состояния юзера
	err := redis.SetStatus(redisClient, services.GetUserID(context), newStatus)
	if err != nil {
		return context.Send(errorMessage, keyboards.BackToHomeInlineKeyboard)
	}

	msgText := `Введите дату в формате ДД/ММ/ГГГГ.
Например, для 23 июля 2016 года вы должны ввести 23/07/2016

❗️История курса валюты доступна с ` + settings.LowDate + " по текущий день"

	return context.Send(msgText)
}

// обработка введённой юзером даты
func ChooseDate(context telebot.Context) error {
	if statusNotIs(context, "choose_date") {
		return nil
	}

	// читаем сообщение юзера (там должна быть дата)
	date := context.Message().Text

	// проверка даты на валидность
	err := services.CheckDate(date)
	// если дата не прошла проверку, то просим юзера ввести её ещё раз
	if err != nil {
		return context.Send(invalidDateMessage)
	}

	newStatus := "gotten_historical_course"
	// установка состояния юзера
	err = redis.SetStatus(redisClient, services.GetUserID(context), newStatus)
	if err != nil {
		return context.Send(errorMessage, keyboards.BackToHomeInlineKeyboard)
	}

	// получение выбранной юзером валюты из redis
	chosenCurrencyCode, err := redis.GetChosenCurrency(redisClient, services.GetUserID(context))
	if err != nil {
		return context.Send(errorMessage, keyboards.BackToHomeInlineKeyboard)
	}

	// получение курса определённой даты из API
	dateCurse, err := currency_api.GetDateCourse(chosenCurrencyCode, date)
	if err != nil {
		return context.Send(notFoundCurrencyMessage, keyboards.BackToHomeInlineKeyboard)
	}

	msgText := fmt.Sprintf("Курс %s за %s - %.2f ₽\n\nВыберите действие 👇", chosenCurrencyCode, date, dateCurse)
	return context.Send(msgText, keyboards.GottenCourseInlineKeyboard)
}
