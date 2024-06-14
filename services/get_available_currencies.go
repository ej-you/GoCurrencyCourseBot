package services

import (
	"encoding/json"
	"os"

	"github.com/Danil-114195722/GoCurrencyCourseBot/settings"
)


type AvailableCurrency struct {
	Code	string	`json:"code"`
	Title	string	`json:"title"`
}


// получение доступных валют из JSON-файла
func GetAvailableCurrencies() ([]AvailableCurrency, error) {
	var curList []AvailableCurrency

	// открытие файла
	fileData, err := os.ReadFile(settings.PathToAvailableCurrencies)
	if err != nil {
		settings.ErrorLog.Println("Failed to open json-file with currency list:", err)
		return curList, err
	}

	// перевод данных из JSON-списка в список структур AvailableCurrency
	err = json.Unmarshal(fileData, &curList)
	if err != nil {
		settings.ErrorLog.Println("Failed to parse json with currency list from file:", err)
		return curList, err
	}

	return curList, nil
}
