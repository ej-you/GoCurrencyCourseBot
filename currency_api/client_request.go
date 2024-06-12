package currency_api

import (
	"net/http"

	"github.com/Danil-114195722/GoCurrencyCourseBot/settings"
)


// Обработка GET-запроса. Возвращает ответ полностью
func getRequest(url string) (*http.Response, error) {
	var emptyResponse *http.Response

	// создание HTTP-клиента
	client := &http.Client{}

	// создание запроса с добавлением заголовка User-Agent
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		// лог с описанием ошибки
		settings.ErrorLog.Println("Failed to create request:", err)
		return emptyResponse, err
	}
	request.Header.Set("User-Agent", "No User-Agent")

	// выполнение запроса
	response, err := client.Do(request)
	if err != nil {
		settings.ErrorLog.Println("Failed to get data from API of Center Bank of Russia:", err)
		return emptyResponse, err
	}

	return response, nil
}