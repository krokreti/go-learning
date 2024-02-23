package main

import (
	"fmt"

	"example.com/structs/user"
)


func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userbirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser *user.User
	
	appUser, err := &user.New(userFirstName, userLastName, userbirthdate)

	if err != nil {
		fmt.Print(err)
		return
	}

	admin := user.NewAdmin("teste@email.com", "teste123")
	
	admin.User.OutputUserDetails()

	appUser.OutputUserDetails()
	appUser.ClearUserName()
	appUser.OutputUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
