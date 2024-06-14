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
}
