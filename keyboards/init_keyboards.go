package keyboards


func InitKeyboards() {
	// клавиатура для команд /start, /home и /cancel
	// кнопки для перехода к командам /course и /currencies
	StartHomeCancelInlineKeyboard.Inline(
		StartHomeCancelInlineKeyboard.Row(BtnCurrrencyCourse),
		StartHomeCancelInlineKeyboard.Row(BtnCurrrencies),
	)

	// клавиатура для команд /currencies, /help
	// кнопка для перехода к команде /home
	BackToHomeInlineKeyboard.Inline(
		BackToHomeInlineKeyboard.Row(BtnBackToHome),
	)

	// клавиатура для команды /course
	// кнопки для выбора валюты
	CurrenciesListInlineKeyboard.Inline(currenciesListBtnRows()...)

	// клавиатура для команды /course
	// кнопки для выбора действия
	ActionInlineKeyboard.Inline(
		ActionInlineKeyboard.Row(BtnActualCourse),
		ActionInlineKeyboard.Row(BtnHistoricalCourse),
	)

	// клавиатура для команды /course
	// кнопки для выбора действия после выдачи курса валюты
	GottenCourseInlineKeyboard.Inline(
		ActionInlineKeyboard.Row(BtnGetCourseAgain),
		ActionInlineKeyboard.Row(BtnCourseBackToHome),
	)
}
