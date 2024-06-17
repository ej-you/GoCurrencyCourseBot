package keyboards

import (
	telebot "gopkg.in/telebot.v3"
)


var BackToHomeInlineKeyboard = &telebot.ReplyMarkup{}
var BtnBackToHome = BackToHomeInlineKeyboard.Data("главное меню", "back_to_home")
