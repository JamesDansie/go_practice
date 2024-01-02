package main

import (
	"fmt"
	"sync"
	"time"
)

const conferenceTickets = 50

var conferenceName string = "Go Conference"
var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {
	greetUsers()

	userFirstName, userLastName, userTickets := getUserInput()

	isValidName, isValidTicketNumber := validateUserInput(userFirstName, userLastName, userTickets)

	if isValidName && isValidTicketNumber {
		bookTicket(userTickets, userFirstName, userLastName)

		wg.Add(1)
		go sendTicket(userTickets, userFirstName, userLastName)

		firstNames := getFirstNames()
		fmt.Printf("These are all our bookings: %v.\n", firstNames)
		if remainingTickets < 1 {
			fmt.Printf("%v is fully booked. Come back next year!\n", conferenceName)
		}
	} else {
		if !isValidName {
			fmt.Println("Invalid user name")
		}
		if !isValidTicketNumber {
			fmt.Println("Invalid ticket number")
		}
	}
	wg.Wait()
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v remaining.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
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
	var userData = UserData{
		firstName:       userFirstName,
		lastName:        userLastName,
		numberOfTickets: userTickets,
	}
	bookings = append(bookings, userData)

	fmt.Printf("User %v %v booked %v tickets.\n", userFirstName, userLastName, userTickets)
	fmt.Printf("%v tickets remaining for %v.\n", remainingTickets, conferenceName)
}

func validateUserInput(userFirstName string, userLastName string, userTickets uint) (bool, bool) {
	isValidName := len(userFirstName) >= 2 && len(userLastName) >= 2
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidTicketNumber
}

func sendTicket(userTickets uint, firstName string, lastName string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("#################")
	fmt.Printf("Sending ticket: %v\n", ticket)
	fmt.Println("#################")
	wg.Done()
}
