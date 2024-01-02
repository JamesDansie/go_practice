package main

import (
	"fmt"
	"strings"
)

const conferenceTickets = 50

var conferenceName string = "Go Conference"
var remainingTickets uint = 50
var bookings []string

func main() {
	greetUsers()

	for remainingTickets > 0 && len(bookings) < 50 {
		userFirstName, userLastName, userTickets := getUserInput()

		isValidName, isValidTicketNumber := validateUserInput(userFirstName, userLastName, userTickets)

		if isValidName && isValidTicketNumber {
			bookTicket(userTickets, userFirstName, userLastName)

			firstNames := getFirstNames()
			fmt.Printf("These are all our bookings: %v.\n", firstNames)
			if remainingTickets < 1 {
				fmt.Printf("%v is fully booked. Come back next year!\n", conferenceName)
				break
			}
		} else {
			if !isValidName {
				fmt.Println("Invalid user name")
			}
			if !isValidTicketNumber {
				fmt.Println("Invalid ticket number")
			}
			continue
		}
	}
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v remaining.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		var firstName = names[0]
		firstNames = append(firstNames, firstName)
	}
	return firstNames
}

func validateUserInput(userFirstName string, userLastName string, userTickets uint) (bool, bool) {
	isValidName := len(userFirstName) >= 2 && len(userLastName) >= 2
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidTicketNumber
}

func getUserInput() (string, string, uint) {
	var userFirstName string
	var userLastName string
	var userTickets uint

	fmt.Println("What is your first name?")
	fmt.Scan(&userFirstName)

	fmt.Println("What is your last name?")
	fmt.Scan(&userLastName)

	fmt.Println("How many tickets would you like?")
	fmt.Scan(&userTickets)
	return userFirstName, userLastName, userTickets
}

func bookTicket(userTickets uint, userFirstName string, userLastName string) {
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, userFirstName+" "+userLastName)

	fmt.Printf("User %v %v booked %v tickets.\n", userFirstName, userLastName, userTickets)
	fmt.Printf("%v tickets remaining for %v.\n", remainingTickets, conferenceName)
}
