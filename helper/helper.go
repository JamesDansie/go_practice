package helper

func ValidateUserInput(userFirstName string, userLastName string, userTickets uint, remainingTickets uint) (bool, bool) {
	isValidName := len(userFirstName) >= 2 && len(userLastName) >= 2
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidTicketNumber
}
