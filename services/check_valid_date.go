package services

import (
	"errors"
	"time"

	"github.com/Danil-114195722/GoCurrencyCourseBot/settings"
)


// проверка, что дата валидна и находится в диапазоне от 01/01/1996 по текущий день
func CheckDate(strDate string) error {
	// верхний предел даты
	currrentDate := time.Now()
	// нижний предел даты
	lowLimitDate, _ := time.Parse("02/01/2006", "01/01/1996")

	// dd/mm/yyyy
	date, err := time.Parse("02/01/2006", strDate)
	if err != nil {
		// лог с описанием ошибки
		settings.ErrorLog.Printf("Failed to parse date from string %q: %v", strDate, err)
		return err
	}

	// если запрашиваемая дата больше текущей
	if date.After(currrentDate) {
		DateMaxError := errors.New("given date more then current date")
		settings.ErrorLog.Printf("Invalid user's date %q: %v", strDate, DateMaxError)
		return DateMaxError
	}

	// если запрашиваемая дата меньше, чем 01/01/1996
	if date.Before(lowLimitDate) {
		DateMinError := errors.New("given date less then 01/01/1996 date")
		settings.ErrorLog.Printf("Invalid user's date %q: %v", strDate, DateMinError)
		return DateMinError
	}

	// если всё хорошо
	return nil
}
