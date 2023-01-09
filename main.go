package main

import (
	"fmt"
	"strings"
)

func main() {

	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	var bookings []string

	fmt.Printf("Welcome to %v booking application !!!\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still there\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	for {

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

		if userTicket > remainingTickets {
			fmt.Printf("we have only %v tickets remaining\n", remainingTickets)
			continue
		}

		remainingTickets = remainingTickets - userTicket
		bookings = append(bookings, userFirstName+" "+userLastName)

		// ask user for name
		fmt.Printf("Thanks you %v %v for booking %v tiskets, you will receive confirmation on your email %v \n", userFirstName, userLastName, userTicket, email)
		fmt.Printf("At %v are tickets %v remaining \n", conferenceName, remainingTickets)
		fmt.Printf("There are all bookings %v \n", bookings)

		firstNames := []string{}
		for _, booking := range bookings {
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}
		fmt.Printf("The first name are for bookings %v \n", firstNames)

		if remainingTickets == 0 {
			// end program
			fmt.Println("Our conference is book out ! See you next year")
			break
		}
	}

}
