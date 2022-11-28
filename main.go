package main

import (
	"fmt"
	"strings"
)

func main() {

	const conferenceTickets int = 50
	var remainingTickets uint = 50
	var bookings = []string{}
	// bookings := []string{}

	fmt.Printf("Welcome to Golang conference\n")
	fmt.Printf("We have %v available tickets & %v remaining tickets\n", conferenceTickets, remainingTickets)

	var firstName string
	var lastName string
	var email string
	var bookedTickets uint
	// Add first and last name to Bookings array

	for {
		fmt.Println("--------------- Book Tickets ---------------")
		// Get firstname
		fmt.Println("Please enter your first name:")
		fmt.Scan(&firstName)

		// Get lasttname
		fmt.Println("Please enter your last name:")
		fmt.Scan(&lastName)

		// Get email
		fmt.Println("Please enter your email:")
		fmt.Scan(&email)

		// Get tickets for bookking
		fmt.Println("How many tickets you want:")
		fmt.Scan(&bookedTickets)

		if bookedTickets <= remainingTickets {

			bookings = append(bookings, firstName+" "+lastName)
			remainingTickets = remainingTickets - uint(bookedTickets)

			// Get first name
			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}

			if remainingTickets == 0 {
				fmt.Println("--------------- Details for Administration ---------------  ")
				fmt.Printf("All Tickets are booked for conference.\n")
				fmt.Printf("Total Bookings: %v \n", firstNames)
				break
			}

			fmt.Printf("Congratulations %v %v 🎉. You have booked %v tickets. A conformation email will be sent at %v.\n",
				firstName, lastName, bookedTickets, email)

			fmt.Println("--------------- Details for Administration ---------------  ")
			fmt.Printf("%v has booked %v tickets. Now only %v tickets are available.\n", bookings[0], bookedTickets, remainingTickets)
			fmt.Printf("Total Bookings: %v \n", firstNames)
		} else {
			fmt.Printf("Sorry, we only have %v available tickets.Try again! \n", remainingTickets)
		}

	}

}
