package main

import "fmt"

func main() {

	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50

	fmt.Printf("Welcome to %v booking application !!!\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still there\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var userFirstName string
	var userLastName string
	var email string
	var userTicket uint

	//userName = "Miloslav"
	fmt.Print("Provide your first name: ")
	fmt.Scan(&userFirstName)
	fmt.Print("Provide your last name: ")
	fmt.Scan(&userLastName)
	fmt.Print("Provide your email address: ")
	fmt.Scan(&email)
	fmt.Print("Provide count of your ticket: ")
	fmt.Scan(&userTicket)

	remainingTickets = remainingTickets - userTicket

	// ask user for name
	fmt.Printf("Thanks you %v %v for booking %v tiskets, you will receive confirmation on your email %v\n Remaining tickets are %v\n", userFirstName, userLastName, userTicket, email, remainingTickets)
}
