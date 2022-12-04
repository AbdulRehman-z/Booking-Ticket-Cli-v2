package helper

import (
	"Golang/models/user"
	"fmt"
	"strings"
)

func GreetUsers(conferenceTickets uint, remainingTickets uint, conferenceName string) {

	fmt.Printf("Welcome to %v Golang conference\n", conferenceName)
	fmt.Printf("We have %v available tickets & %v remaining tickets for %v Golang conference. \n", conferenceTickets, remainingTickets, conferenceName)
}

func ValidateUserInputs(firstName string, lastName string, email string, bookedTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidUsername := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidBooking := bookedTickets <= remainingTickets

	return isValidUsername, isValidEmail, isValidBooking
}

func GetFirstNames(bookings []string) []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames

}

func GetUserInputs() user.UserInputs {
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

	userData := user.UserInputs{
		FirstName:     firstName,
		LastName:      lastName,
		Email:         email,
		BookedTickets: bookedTickets,
	}

	return userData
}

func ErrorReport(isValidUsername bool, isValidEmail bool, isValidBooking bool, firstName string, email string, remainingTickets uint) {

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

func ZeroRemainingTickets(firstNames []string) {
	fmt.Println("--------------- Details for Administration ---------------  ")
	fmt.Printf("All Tickets are booked for conference.\n")
	fmt.Printf("Total Bookings: %v \n", firstNames)
}

func AdministrationDetails(firstNames []string, bookings []string, bookedTickets uint, remainingTickets uint) {
	fmt.Println("--------------- Details for Administration ---------------  ")
	fmt.Printf("%v has booked %v tickets. Now only %v tickets are available.\n", bookings[0], bookedTickets, remainingTickets)
	fmt.Printf("Total Bookings: %v \n", firstNames)

}
