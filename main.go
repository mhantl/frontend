package main

import (
	"booking-app/helper"
	"fmt"
	"strings"
)

var conferenceName = "Go Conference"

const conferenceTickets int = 50

var remainingTickets uint = 50
var bookings []string

//var isValidName, isValidEmail, isValidTicketNumber bool

func main() {

	greetUsers()

	for {

		firstName, lastName, email, userTickets := getUserInput()

		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			if userTickets > remainingTickets {
				fmt.Printf("we have only %v tickets remaining\n", remainingTickets)
				continue
			}

			// call func booking
			bookTicket(firstName, lastName, email, userTickets)

			// call func first name
			var firstNames = getFirstNames()
			fmt.Printf("The first name are for bookings %v \n", firstNames)

			if remainingTickets == 0 {
				// end program
				fmt.Println("Our conference is book out ! See you next year")
				break
			}
		} else {
			if !isValidName {
				fmt.Print("First or last name is not valid - is too short\n")
			}
			if !isValidEmail {
				fmt.Print("Email is not valid - not contains @ \n")
			}
			if !isValidTicketNumber {
				fmt.Print("number of tickets you entered is invalid\n")
			}
		}
	}

}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application \n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still there\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}

	return firstNames

}

func getUserInput() (string, string, string, uint) {
	//userName = "Miloslav"

	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Print("Provide your first name: ")
	fmt.Scan(&firstName)
	fmt.Print("Provide your last name: ")
	fmt.Scan(&lastName)
	fmt.Print("Provide your email address: ")
	fmt.Scan(&email)
	fmt.Print("Provide count of your ticket: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(firstName string, lastName string, email string, userTickets uint) {
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName+" "+lastName)

	// ask user for name
	fmt.Printf("Thanks you %v %v for booking %v tiskets, you will receive confirmation on your email %v \n", firstName, lastName, userTickets, email)
	fmt.Printf("At %v are tickets %v remaining \n", conferenceName, remainingTickets)
	fmt.Printf("There are all bookings %v \n", bookings)
}
