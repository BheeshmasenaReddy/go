package main

import (
	"fmt"
	"strings"
)

const conferenceTickets = 50

var conferenceName = "Bheeshma's Conference"
var remainingTickets uint = 50
var bookings = []string{}

func main() {

	//var str = "Hello World! This is my first program in go!! Let's Goooo!!"
	greetUser()

	for {
		var firstName, lastName, email, tickets = getUserInput()

		var isValidName, isValiEmail, isValidTickets = validateUserInput(firstName, lastName, email, tickets)

		if isValidName && isValiEmail && isValidTickets {

			bookTicket(firstName, lastName, email, tickets)
			var firstNames = getFirstNames()

			fmt.Printf("The remaining number of tickets are: %v\n", remainingTickets)
			fmt.Printf("These are the people who made bookings till now: %v\n", firstNames)

			fmt.Printf("Total Number of bookings: %v\n", len(bookings))
			fmt.Printf("The first user to book tickets is: %v\n", bookings[0])

			fmt.Println("Thanks for booking the tickets for", conferenceName, "see you there!")

			if remainingTickets == 0 {
				fmt.Println("The tickets have been sold out. Please do come again next year!! Thank You!")
				break
			}

		} else {
			if !isValidName {
				fmt.Println("Please enter your name correct! It's too short.")
			}
			if !isValiEmail {
				fmt.Println("Your email address is not valid because it does not contain '@' sign")
			}

			if !isValidTickets {
				fmt.Printf("Sorry! You can't book %v tickets, we only have %v tickets\n", tickets, remainingTickets)
				fmt.Println("Please try again to book available tickets!!")
			}

		}

	}
}

func greetUser() {
	fmt.Printf("Welcome to our %v booking application!!\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available for booking\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var tickets uint

	fmt.Print("Enter your first name: \n")
	fmt.Scan(&firstName)

	fmt.Print("Enter your last name: \n")
	fmt.Scan(&lastName)

	fmt.Print("Enter your email address: \n")
	fmt.Scan(&email)

	fmt.Print("Enter the number of tickets you want to book: \n")
	fmt.Scan(&tickets)

	return firstName, lastName, email, tickets
}

func validateUserInput(firstName string, lastName string, email string, tickets uint) (bool, bool, bool) {
	isValidName := len(firstName) > 2 && len(lastName) > 2
	isValiEmail := strings.Contains(email, "@")
	isValidTickets := tickets > 0 && tickets <= remainingTickets

	return isValidName, isValiEmail, isValidTickets
}

func bookTicket(firstName string, lastName string, email string, tickets uint) {
	bookings = append(bookings, firstName+" "+lastName)
	remainingTickets -= tickets
	fmt.Printf("%v booked %v tickets, You will get confirmation to your mail %v \n", firstName+" "+lastName, tickets, email)

}

func getFirstNames() []string {
	var firstNames []string
	for _, name := range bookings {
		firstNames = append(firstNames, strings.Fields(name)[0])
	}
	return firstNames
}
