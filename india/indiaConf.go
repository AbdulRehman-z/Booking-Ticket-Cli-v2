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
		firstName, lastName, email, bookedTickets := helper.GetUserInputs()

		// Validate User Inputs
		isValidUsername, isValidEmail, isValidBooking := helper.ValidateUserInputs(firstName, lastName, email, bookedTickets, remainingTickets)

		if isValidUsername && isValidEmail && isValidBooking {

			bookings = append(bookings, firstName+" "+lastName)
			remainingTickets = remainingTickets - bookedTickets

			firstNames := helper.GetFirstNames(bookings)

			fmt.Printf("Congratulations %v %v 🎉. You have booked %v tickets. A conformation email will be sent at %v.\n",
				firstName, lastName, bookedTickets, email)

			// Details for administration about the bookings
			helper.AdministrationDetails(firstNames, bookings, bookedTickets, remainingTickets)
			// Check if remainingTickets == 0
			if remainingTickets == 0 {
				helper.ZeroRemainingTickets(firstNames)
				break
			}
		} else {
			helper.ErrorReport(isValidUsername, isValidEmail, isValidBooking, firstName, email, remainingTickets)
		}
	}
}
