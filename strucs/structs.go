package main

import (
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName string
	birthDate string
	createdAt time.Time
}

func main() {
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	// ... do something awesome with that gathered data!

	outputUserDetails(firstName, lastName, birthdate)
}

func outputUserDetails(firstname, lastName, birthdate string) {
	fmt.Print(firstname, lastName, birthdate)
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
