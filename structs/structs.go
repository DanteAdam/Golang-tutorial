package main

import (
	"fmt"
	"time"
)

type user struct {
	firstName, lastName, birthDate string
	createAt                       time.Time
}

func main() {
	userfirstName := getUserData("Please enter your first name:")
	userlastName := getUserData("Please enter your last name:")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY):")

	var appUser user

	appUser = user{
		firstName: userfirstName,
		lastName:  userlastName,
		birthDate: userBirthdate,
		createAt:  time.Now(),
	}
	ouputUserDetails(appUser)
}
func ouputUserDetails(u user) {
	fmt.Println(u.firstName, u.lastName, u.birthDate)

}
func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
