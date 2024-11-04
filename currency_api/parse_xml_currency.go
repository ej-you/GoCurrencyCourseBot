package currency_api

import (
	"encoding/xml"
	"errors"
	"net/http"

	"golang.org/x/text/encoding/charmap"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/transform"

	"github.com/ej-you/GoCurrencyCourseBot/settings"
)


// Парсинг всех валют, полученных из XML. Возвращает структуру CurrencyList
func parseCurrencyList(response *http.Response) (CurrencyList, error) {
	var resList CurrencyList

	// перекодировка из windows-1251 в utf-8
	decoder := charmap.Windows1251.NewDecoder()
	// создание XML-декодера, который сможет обрабатывать windows-1251
	xmlDecoder := xml.NewDecoder(transform.NewReader(response.Body, decoder))
	// установка функции для обработки кодировки
	xmlDecoder.CharsetReader = charset.NewReaderLabel

	// парсинг XML данных
	err := xmlDecoder.Decode(&resList)
	if err != nil {
		settings.ErrorLog.Println("Failed to parse XML:", err)
		return resList, err
	}

	return resList, nil
}

// Получение инфо о конкретной валюте из списка валют, полученных из XML. Возвращает структуру Currency
func getCurrency(response *http.Response, currencyCode string) (Currency, error) {
	var emptyCurrency Currency

	// получение списка валют из XML в виде структуры CurrencyList
	fullList, err := parseCurrencyList(response)
	if err != nil {
		return emptyCurrency, err
	}

	// поиск нужной валюты по всем спарсенным
	for _, currency := range(fullList.Currencies) {
		if currency.Code == currencyCode {
			// если вылюта была найдена по коду
			return currency, nil
		}
	}
	// если валюта не найдена
	notFoundError := errors.New("given currency was not found in parsed XML")
	settings.ErrorLog.Printf("Currency %q not found in XML: %v", currencyCode, notFoundError)
	return emptyCurrency, notFoundError
}