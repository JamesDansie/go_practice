package main

import (
	"fmt"
	"strings"
)

func main() {
	// This method replaces the type and var declaration
	conferenceName := "Go Conference"
	const conferenceTickets = 50
	// overriding type to make sure we never have a negative amount of tickets
	var remainingTickets uint = 50

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v remaining.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var bookings []string
	for {

		var userFirstName string
		var userLastName string
		var userTickets uint

		fmt.Println("What is your first name?")
		fmt.Scan(&userFirstName)

		fmt.Println("What is your last name?")
		fmt.Scan(&userLastName)

		fmt.Println("How many tickets would you like?")
		fmt.Scan(&userTickets)

		remainingTickets = remainingTickets - userTickets
		bookings = append(bookings, userFirstName+" "+userLastName)

		fmt.Printf("User %v %v booked %v tickets.\n", userFirstName, userLastName, userTickets)
		fmt.Printf("%v tickets remaining for %v.\n", remainingTickets, conferenceName)

		firstNames := []string{}
		for _, booking := range bookings {
			var names = strings.Fields(booking)
			var firstName = names[0]
			firstNames = append(firstNames, firstName)
		}
		fmt.Printf("These are all our bookings: %v.\n", firstNames)

		if remainingTickets < 1 {
			fmt.Printf("%v is fully booked. Come back next year!\n", conferenceName)
			break
		}
	}
}
