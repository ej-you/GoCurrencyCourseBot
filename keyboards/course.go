package keyboards

import (
	telebot "gopkg.in/telebot.v3"
	
	"github.com/ej-you/GoCurrencyCourseBot/services"
)

// Клавиатура с доступными валютами
var CurrenciesListInlineKeyboard = &telebot.ReplyMarkup{}
// Возвращает ряды кнопок с доступными валютами (по 3 кнопки в каждом ряду)
func currenciesListBtnRows() []telebot.Row {
	// переменная для добавления каждой новой кнопоки
	var nextButton telebot.Btn
	// срез кнопок с доступными валютами
	var curButtons []telebot.Btn
	// получение доступных валют из JSON-файла
	curList, _ := services.GetAvailableCurrencies()

	for _, cur := range curList {
		nextButton = CurrenciesListInlineKeyboard.Data(cur.Code, cur.Code)
		curButtons = append(curButtons, nextButton)
	}

	// переменная для добавления каждого нового ряда кнопок
	var nextRow telebot.Row
	// срез для добавления рядов кнопок
	var inlineRows []telebot.Row
	for i := 0; i <= (len(curButtons) / 3 * 3); i+=3 {
		if i + 3 > len(curButtons) {
			nextRow = CurrenciesListInlineKeyboard.Row(curButtons[i:]...)
		} else {
			nextRow = CurrenciesListInlineKeyboard.Row(curButtons[i:i+3]...)
		}
		inlineRows = append(inlineRows, nextRow)
	}
	return inlineRows
}

var ActionInlineKeyboard = &telebot.ReplyMarkup{}
var BtnActualCourse = ActionInlineKeyboard.Data("актуальный курс", "actual_course")
var BtnHistoricalCourse = ActionInlineKeyboard.Data("курс за определённую дату", "historical_course")

var GottenCourseInlineKeyboard = &telebot.ReplyMarkup{}
var BtnGetCourseAgain = GottenCourseInlineKeyboard.Data("запросить ещё", "get_course_again")
var BtnCourseBackToHome = GottenCourseInlineKeyboard.Data("главное меню", "course_back_to_home")
