package keyboards

import (
	telebot "gopkg.in/telebot.v3"
)


var StartHomeCancelInlineKeyboard = &telebot.ReplyMarkup{}
var BtnCurrrencyCourse = StartHomeCancelInlineKeyboard.Data("узнать курс валюты", "get_cur_course")
var BtnCurrrencies = StartHomeCancelInlineKeyboard.Data("список валют", "get_currencies")
