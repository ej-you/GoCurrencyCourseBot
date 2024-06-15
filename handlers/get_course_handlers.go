package handlers

import (
	"fmt"
	"strconv"

	telebot "gopkg.in/telebot.v3"

	"github.com/Danil-114195722/GoCurrencyCourseBot/currency_api"
	"github.com/Danil-114195722/GoCurrencyCourseBot/keyboards"
	"github.com/Danil-114195722/GoCurrencyCourseBot/redis"
	"github.com/Danil-114195722/GoCurrencyCourseBot/services"
)


// тестовый запуск функций получения курса валюты
// getCurrencyCourseTest("USD", "19/11/2006")
func getCurrencyCourseTest(currencyCode, date string) {
	latestCurse, err := currency_api.GetLatestCourse(currencyCode)

	if err == nil {
		fmt.Printf("Course %s now: %.4f\n", currencyCode, latestCurse)
	} else {
		fmt.Println("NOT FOUND!!!")
	}

	err = services.CheckDate(date)
	if err != nil {
		fmt.Println("INVALID DATE!!!")
	} else {
		dateCurse, err := currency_api.GetDateCourse(currencyCode, date)

		if err == nil {
			fmt.Printf("Course %s at %s: %.4f\n", currencyCode, date, dateCurse)
		} else {
			fmt.Println("NOT FOUND!!!")
		}
	}
}

// команда /course
func CourseHandler(context telebot.Context) error {
	// строковый id юзера ТГ
	userIdString := strconv.FormatInt(context.Chat().ID, 10)

	// установка состояния юзера
	err := redis.SetStatus(redisClient, userIdString, "start")
	if err != nil {
		return context.Send(errorMessage, keyboards.BackToHomeInlineKeyboard)
	}

	// получение состояния юзера
	status, err := redis.GetStatus(redisClient, userIdString)
	if err != nil {
		return context.Send(errorMessage, keyboards.BackToHomeInlineKeyboard)
	}

	msgText := "course handler: " + status
	return context.Send(msgText)
}
