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
	fmt.Printf("conferenceTickets is %T\n", conferenceTickets)

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v remaining.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var userName string
	var userTickets uint
	// ask user for their name
	fmt.Println("What is your name?")
	fmt.Scan(&userName)
	fmt.Println("How many tickets would you like?")
	fmt.Scan(&userTickets)
	fmt.Printf("User %v booked %v tickets.\n", userName, userTickets)

	// print vs print pointer
	fmt.Println(remainingTickets)
	fmt.Println(&remainingTickets)
}
