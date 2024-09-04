package main

import (
	"fmt"
	"time"
)

type User struct {
	firstName, lastName, birthDate string 
	create time.Time
}

func main() {
	firstName := getUserData("Please enter your first name :")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser user 

	appUser = 

	fmt.Println(firstName, lastName, birthdate)
}
func ouputUserDetails(){
	fmt.Println(firstName, lastName, birthdate string){

	}
}
func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
