package main

import (
	"fmt"
	"strings"
)

// Prints the given string to stdout
func main() {
	var conferenceName = "Bheeshma's Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	var bookings = []string{}

	fmt.Printf("Welcome to our %v booking application!!\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available for booking\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
	//var str = "Hello World! This is my first program in go!! Let's Goooo!!"

	for {
		var firstName string
		var lastName string
		var tickets uint

		fmt.Print("Enter your first name: \n")
		fmt.Scan(&firstName)

		fmt.Print("Enter your last name: \n")
		fmt.Scan(&lastName)

		fmt.Print("Enter the number of tickets you want to book: \n")
		fmt.Scan(&tickets)

		fmt.Printf("%v booked %v tickets\n", firstName+" "+lastName, tickets)

		bookings = append(bookings, firstName+" "+lastName)

		remainingTickets -= tickets
		var firstNames []string
		for _, name := range bookings {
			firstNames = append(firstNames, strings.Fields(name)[0])
		}

		fmt.Printf("The remaining number of tickets are: %v\n", remainingTickets)
		fmt.Printf("These are the people who made bookings till now: %v\n", firstNames)

		fmt.Printf("Total Number of bookings: %v\n", len(bookings))
		fmt.Printf("The first user to book tickets is: %v\n", bookings[0])
		fmt.Printf("Type of bookings is: %T\n", bookings)

		fmt.Println("Thanks for booking the tickets for", conferenceName, "see you there!")
	}
}
