package pakistan

import (
	"Golang/helper"
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func PakistanConference() {
	const conferenceTickets uint = 40
	var remainingTickets uint = 40
	conferenceName := "Tokoyo"
	bookings := []string{}

	// Greet users to the conference
	helper.GreetUsers(conferenceTickets, remainingTickets, conferenceName)

	// Get user Inputs
	userData := helper.GetUserInputs()
	// Validate User Inputs
	isValidUsername, isValidEmail, isValidBooking := helper.ValidateUserInputs(userData.FirstName, userData.LastName, userData.Email, userData.BookedTickets, remainingTickets)

	if isValidUsername && isValidEmail && isValidBooking {

		bookings = append(bookings, userData.FirstName)
		remainingTickets = remainingTickets - userData.BookedTickets
		wg.Add(1)
		go sendTicket(userData.FirstName)

		firstNames := helper.GetFirstNames(bookings)
		fmt.Printf("Congratulations %v %v 🎉. You have booked %v tickets. A conformation email will be sent at %v.\n",
			userData.FirstName, userData.LastName, userData.BookedTickets, userData.Email)

		// Details for administration about the bookings
		helper.AdministrationDetails(firstNames, bookings, userData.BookedTickets, remainingTickets)
		// Check if remainingTickets == 0
		if remainingTickets == 0 {
			helper.ZeroRemainingTickets(firstNames)
		}
	} else {
		helper.ErrorReport(isValidUsername, isValidEmail, isValidBooking, userData.FirstName, userData.Email, remainingTickets)
	}

	wg.Wait()
}

func sendTicket(firstName string) {
	time.Sleep(20 * time.Second)
	ticket := fmt.Sprintf("%v booked ticket!", firstName)
	fmt.Println("##################")
	fmt.Printf("%v\n", ticket)
	fmt.Println("##################")
	wg.Done()
}
