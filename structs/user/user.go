package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName, lastName, birthDate string
	createAt                       time.Time
}

type Admin struct {
	email, password string
	User
}

func (u *User) OuputUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthDate)

}
func (u *User) ClearUserName() {
	u.firstName = ""
	u.lastName = ""

}
func NewAdmin(email, password string) *Admin {
	return &Admin{
		email:    email,
		password: password,
		User: User{
			firstName: "ADMIN",
			lastName:  "ADMIN",
			birthDate: "----",
			createAt:  time.Now(),
		},
	}
}
func New(firstName, lastName, birthdate string) (*User, error) {
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("first name, last name and birthdate are required")
	}

	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthdate,
		createAt:  time.Now(),
	}, nil

}
