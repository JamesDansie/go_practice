package main

import (
	"fmt"
)

func main() {
	// This method replaces the type and var declaration
	conferenceName := "Go Conference"
	const conferenceTickets = 50
	// overriding type to make sure we never have a negative amount of tickets
	var remainingTickets uint = 50

	// handy for debugging types
	// fmt.Printf("conferenceTickets is %T\n", conferenceTickets)

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v remaining.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var bookings [50]string

	var userFirstName string
	var userLastName string
	var userTickets uint
	// ask user for their name
	fmt.Println("What is your first name?")
	fmt.Scan(&userFirstName)
	fmt.Println("What is your last name?")
	fmt.Scan(&userLastName)
	fmt.Println("How many tickets would you like?")
	fmt.Scan(&userTickets)
	remainingTickets = remainingTickets - userTickets
	bookings[0] = userFirstName + " " + userLastName
	fmt.Printf("bookings: %v\n", bookings)
	fmt.Printf("bookings[0]: %v\n", bookings[0])
	fmt.Printf("bookings type: %T\n", bookings)
	fmt.Printf("bookings length: %v\n", len(bookings))

	fmt.Printf("User %v %v booked %v tickets.\n", userFirstName, userLastName, userTickets)
	fmt.Printf("%v tickets remaining for %v.\n", remainingTickets, conferenceName)

	// print vs print pointer
	// fmt.Println(remainingTickets)
	// fmt.Println(&remainingTickets)
}
