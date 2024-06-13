package currency_api

import (
	"fmt"
	"strings"
	"strconv"

	"github.com/Danil-114195722/GoCurrencyCourseBot/settings"
)


// Возвращает курс запрашиваемой валюты по данной ссылке
func getCourse(currencyCode, apiUrl string) (float64, error) {
	var course float64
	
	// GET-запрос и получение ответа
	response, err := getRequest(apiUrl)
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


// Возвращает текущий курс запрашиваемой валюты
func GetLatestCourse(currencyCode string) (float64, error) {
	// получение курса одной единицы валюты
	course, err := getCourse(currencyCode, "http://www.cbr.ru/scripts/XML_daily.asp/")

	return course, err
}

func GetDateCourse(currencyCode string, strDate string) (float64, error) {
	// URL для запроса курсов валют по данной дате
	urlWithParam := fmt.Sprintf("http://www.cbr.ru/scripts/XML_daily.asp/?date_req=%s", strDate)
	// получение курса одной единицы валюты
	course, err := getCourse(currencyCode, urlWithParam)

	return course, err
}
