package handlers

import (
	"fmt"

	telebot "gopkg.in/telebot.v3"

	"github.com/ej-you/GoCurrencyCourseBot/keyboards"
	"github.com/ej-you/GoCurrencyCourseBot/redis"
	"github.com/ej-you/GoCurrencyCourseBot/services"
)


// команда /start
func StartHandler(context telebot.Context) error {
	newStatus := "home"
	// установка состояния юзера
	err := redis.SetStatus(redisClient, services.GetUserID(context), newStatus)
	if err != nil {
		return context.Send(errorMessage, keyboards.BackToHomeInlineKeyboard)
	}

	msgText := `Привет 👋

Этот бот поможет вам узнать курс популярных валют 📈
Есть возможность узнать не только актуальный курс валюты, но и её курс за определённую дату 🤯

❗️Для получения полной инструкции введите /help`

	return context.Send(msgText, keyboards.StartHomeCancelInlineKeyboard)
}

// команда /home и /cancel
func HomeHandler(context telebot.Context) error {
	newStatus := "home"
	// установка состояния юзера
	err := redis.SetStatus(redisClient, services.GetUserID(context), newStatus)
	if err != nil {
		return context.Send(errorMessage, keyboards.BackToHomeInlineKeyboard)
	}

	msgText := `Вы в главном меню 🗂

❗️Для получения полной инструкции введите /help`

	return context.Send(msgText, keyboards.StartHomeCancelInlineKeyboard)
}

// команда /help
func HelpHandler(context telebot.Context) error {
	msgText := `Полная инструкция и описание этого бота.

Этот бот поможет вам узнать курс популярных валют 📈
Также есть возможность узнать не только актуальный курс валюты, но и её курс за определённую дату 🤯

Команды бота:
		/start - перезапуск бота
		/home - главное меню бота
		/course - активация режима для получения курса валюты
		/currencies - список доступных валют и их кодов
		/help - вывод этой справки
		/cancel - отмена всех действий и переход в главное меню

В режиме получения курса валюты будет несколько этапов.
В каждом шаге подробно описаны требования к действиям пользователя. Просим ВАС их соблюдать.

При запросе исторического значения курса валюты до 2000 года может возникать ошибка.
Информация по некоторым валютам до 2000 года, к сожалению, отсутствует.

Если что-то не получается и вам выдаётся ошибка, просим перечитать требования ещё раз и повторить запрос.`

	return context.Send(msgText, keyboards.BackToHomeInlineKeyboard)
}

// команда /currencies
func CurrenciesHandler(context telebot.Context) error {
	if statusNotIs(context, "home") {
		return nil
	}

	// получение доступных валют из JSON-файла
	curList, err := services.GetAvailableCurrencies()
	if err != nil {
		return context.Send(errorMessage, keyboards.BackToHomeInlineKeyboard)
	}

	msgText := "Список доступных валют и их кодов:\n\nкод: полное название\n-----\n"
	// добавление в текст ответа всех доступных валют
	for _, cur := range curList {
		msgText += fmt.Sprintf("\n%s: %s", cur.Code, cur.Title)
	}
	// ответ бота
	return context.Send(msgText, keyboards.BackToHomeInlineKeyboard)
}
