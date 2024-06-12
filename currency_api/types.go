package currency_api

import (
	"encoding/xml"
)


// Данные одной валюты
type Currency struct {
	Code 		string 		`xml:"CharCode"`
	Nominal		int 		`xml:"Nominal"`
	Value		string		`xml:"Value"`
}

// Список всех распарсенных валют из XML
type CurrencyList struct {
	XMLName		xml.Name	`xml:"ValCurs"`
	Currencies	[]Currency	`xml:"Valute"`
}
