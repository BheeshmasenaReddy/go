package main

import "strings"

func validateUserInput(firstName string, lastName string, email string, tickets uint) (bool, bool, bool) {
	isValidName := len(firstName) > 2 && len(lastName) > 2
	isValiEmail := strings.Contains(email, "@")
	isValidTickets := tickets > 0 && tickets <= remainingTickets

	return isValidName, isValiEmail, isValidTickets
}
