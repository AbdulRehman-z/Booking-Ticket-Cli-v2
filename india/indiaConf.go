package india

import (
	"Golang/helper"
	"fmt"
)

func IndiaConference() {
	const conferenceTickets uint = 40
	var remainingTickets uint = 40
	conferenceName := "Tokoyo"
	bookings := []string{}

	// Greet users to the conference
	helper.GreetUsers(conferenceTickets, remainingTickets, conferenceName)

	for {
		// Get user Inputs
		userData := helper.GetUserInputs()
		// Validate User Inputs
		isValidUsername, isValidEmail, isValidBooking := helper.ValidateUserInputs(userData.FirstName, userData.LastName, userData.Email, userData.BookedTickets, remainingTickets)

		if isValidUsername && isValidEmail && isValidBooking {

			bookings = append(bookings, userData.FirstName)
			remainingTickets = remainingTickets - userData.BookedTickets

			firstNames := helper.GetFirstNames(bookings)
			fmt.Printf("Congratulations %v %v 🎉. You have booked %v tickets. A conformation email will be sent at %v.\n",
				userData.FirstName, userData.LastName, userData.BookedTickets, userData.Email)

			// Details for administration about the bookings
			helper.AdministrationDetails(firstNames, bookings, userData.BookedTickets, remainingTickets)
			// Check if remainingTickets == 0
			if remainingTickets == 0 {
				helper.ZeroRemainingTickets(firstNames)
				break
			}
		} else {
			helper.ErrorReport(isValidUsername, isValidEmail, isValidBooking, userData.FirstName, userData.Email, remainingTickets)
		}
	}
}
