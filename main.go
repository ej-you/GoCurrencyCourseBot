package main

import (
	"fmt"

	"github.com/Danil-114195722/GoCurrencyCourseBot/currency_api"
)

func main() {
	// latestCurse, err := currency_api.GetLatestCourse("USD")
	latestCurse, err := currency_api.GetLatestCourse("SEK")

	if err == nil {
		fmt.Printf("Course: %v\n", latestCurse)
	} else {
		fmt.Println("NOT FOUND!!!")
	}
}
