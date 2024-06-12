package currency_api

import (
	"time"
	"strings"
	"strconv"

	"github.com/Danil-114195722/GoCurrencyCourseBot/settings"
)


// Возвращает текущий курс запрашиваемой валюты
func GetLatestCourse(currencyCode string) (float64, error) {
	var course float64
	
	// GET-запрос и получение ответа
	response, err := getRequest("http://www.cbr.ru/scripts/XML_daily.asp/")
	defer response.Body.Close()
	if err != nil {
		return course, err
	}

	// получение структуры Currency для найденной валюты
	currency, err := getCurrency(response, currencyCode)
	if err != nil {
		return course, err
	}

	// перевод курса номинала из строки во float64
	currencyFloatValue, err := strconv.ParseFloat(strings.ReplaceAll(currency.Value, ",", "."), 64)
	if err != nil {
		settings.ErrorLog.Printf("Failed to parse float64 from string of currency value %q: %v", currency.Value, err)
		return course, err
	}

	// высчитываем реальный курс одной единицы валюты
	course = currencyFloatValue / float64(currency.Nominal)
	return course, nil
}

func GetDateCourse(currencyCode string, date time.Time) float64 {
	return 0.0
}
