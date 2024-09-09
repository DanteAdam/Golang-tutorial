package main

import (
	"fmt"
	"structs/user"
)

func main() {
	userFirstName := getUserData("Please enter your first name:")
	userLastName := getUserData("Please enter your last name:")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY):")

	var appUser, err = user.New(userFirstName, userLastName, userBirthdate)

	if err != nil {
		fmt.Println(err)
		return
	}

	admin := user.NewAdmin("test@example", "123456")
	admin.OuputUserDetails()
	admin.ClearUserName()
	admin.ClearUserName()
	// ouputUserDetails(&appUser)

	appUser.OuputUserDetails()
	appUser.ClearUserName()
	appUser.OuputUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
