package main

import (
	"fmt"
	"strings"
)

func main() {

	const conferenceTickets int = 50
	var remainingTickets uint = 50
	bookings := []string{}

	greetUsers(uint(conferenceTickets), remainingTickets)

	for {
		// Get inputs from user
		firstName, lastName, email, bookedTickets := gerUserInputs()

		// Valide data that user provided
		isValidUsername, isValidEmail, isValidBooking := validateUserInputs(firstName, lastName, email, bookedTickets, remainingTickets)

		if isValidUsername && isValidEmail && isValidBooking {

			bookings = append(bookings, firstName+" "+lastName)
			remainingTickets = remainingTickets - uint(bookedTickets)

			// Get first name
			firstNames := getFirstNames(bookings)

			fmt.Printf("Congratulations %v %v 🎉. You have booked %v tickets. A conformation email will be sent at %v.\n",
				firstName, lastName, bookedTickets, email)

			fmt.Println("--------------- Details for Administration ---------------  ")
			fmt.Printf("%v has booked %v tickets. Now only %v tickets are available.\n", bookings[0], bookedTickets, remainingTickets)
			fmt.Printf("Total Bookings: %v \n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("--------------- Details for Administration ---------------  ")
				fmt.Printf("All Tickets are booked for conference.\n")
				fmt.Printf("Total Bookings: %v \n", firstNames)
				break
			}

		} else {

			if !isValidUsername {
				fmt.Printf("Not valid username: %v. Try again!\n ", firstName)
			}

			if !isValidEmail {
				fmt.Printf("Not valid email address: %v.Missing @ or .com. Try again!\n", email)
			}

			if !isValidBooking {
				fmt.Printf("Sorry, we only have %v available tickets.Try again!\n", remainingTickets)
			}
		}

	}

}

/////////////////////////////////////////
/////////////////////////////////////////
/////////////////////////////////////////
//////// helpers

func greetUsers(conferenceTickets uint, remainingTickets uint) {
	fmt.Printf("Welcome to Golang conference\n")
	fmt.Printf("We have %v available tickets & %v remaining tickets\n", conferenceTickets, remainingTickets)
}

func gerUserInputs() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var bookedTickets uint

	fmt.Println("--------------- Book Tickets -----------------")
	// Get firstname
	fmt.Println("Please enter your first name:")
	fmt.Scan(&firstName)

	// Get lastname
	fmt.Println("Please enter your last name:")
	fmt.Scan(&lastName)

	// Get email
	fmt.Println("Please enter your email:")
	fmt.Scan(&email)

	// Get tickets for bookking
	fmt.Println("How many tickets you want:")
	fmt.Scan(&bookedTickets)

	return firstName, lastName, email, bookedTickets
}

func validateUserInputs(firstName string, lastName string, email string, bookedTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidUsername := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidBooking := bookedTickets <= remainingTickets

	return isValidUsername, isValidEmail, isValidBooking
}

func getFirstNames(bookings []string) []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames

}
