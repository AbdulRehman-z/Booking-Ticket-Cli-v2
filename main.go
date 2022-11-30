package main

import (
	"Golang/germany"
	"Golang/india"
	"Golang/pakistan"
	"Golang/tokoyo"
	"fmt"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func main() {
	var conferenceCity string

	fmt.Println("Welcome to Golang conference. Please enter which conference you want to attend!")
	fmt.Scan(&conferenceCity)
	// Captalize first letter of conferenceCity
	cityName := cases.Title(language.English, cases.Compact).String(conferenceCity)

	switch cityName {
	case "Tokoyo":
		tokoyo.TokoyoConference()
	case "Pakistan":
		pakistan.PakistanConference()
	case "Germany":
		germany.GermanyConference()
	case "India":
		india.IndiaConference()
	default:
		fmt.Println("Provided country or city name does not held any Golang conference. Please Try again!")
	}
}
