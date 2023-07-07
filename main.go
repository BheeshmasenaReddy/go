package main

import "fmt"

// Prints the given string to stdout
func main() {
	var conferenceName = "Bheeshma's Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50

	fmt.Printf("Welcome to our %v booking application!!\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available for booking\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
	//var str = "Hello World! This is my first program in go!! Let's Goooo!!"

	var userName string
	var tickets uint

	fmt.Print("Enter your name: ")
	fmt.Scan(&userName)

	fmt.Print("Enter the number of tickets you want to book: ")
	fmt.Scan(&tickets)

	fmt.Printf("%v booked %v tickets\n", userName, tickets)

	remainingTickets -= tickets

	fmt.Printf("The remaining number of tickets are: %v\n", remainingTickets)

	fmt.Println(conferenceName)

}
