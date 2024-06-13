package main

import (
	"fmt"

	"github.com/Danil-114195722/GoCurrencyCourseBot/currency_api"
	"github.com/Danil-114195722/GoCurrencyCourseBot/services"
)

func main() {
	currencyCode := "EUR"
	date := "15/08/2014"

	latestCurse, err := currency_api.GetLatestCourse(currencyCode)
	// latestCurse, err := currency_api.GetLatestCourse("SEK")

	if err == nil {
		fmt.Printf("Course %s now: %.4f\n", currencyCode, latestCurse)
	} else {
		fmt.Println("NOT FOUND!!!")
	}

	err = services.CheckDate(date)
	if err != nil {
		fmt.Println("INVALID DATE!!!")
	} else {
		dateCurse, err := currency_api.GetDateCourse(currencyCode, date)

		if err == nil {
			fmt.Printf("Course %s at %s: %.4f\n", currencyCode, date, dateCurse)
		} else {
			fmt.Println("NOT FOUND!!!")
		}
	}

}
