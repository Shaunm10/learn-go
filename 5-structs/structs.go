package main

import (
	"errors"
	"fmt"
	"time"
)

func main() {

	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	// ... do something awesome with that gathered data!
	user := User{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthdate,
		createdAt: time.Now(),
	}

	//fmt.Println(firstName, user.lastName, birthdate)
	user.speakMyName()
	user.clearMyName()
	user.speakMyName()

	jack := newUser("Jack", "Sargent", "12/20/2021")

}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}

type User struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

func (user User) speakMyName() {

	fmt.Println(user.firstName, user.lastName, user.birthDate)
}

func (user *User) clearMyName() {
	user.firstName = ""
	user.lastName = ""
}

func newUser(firstName, lastName, birthdate string) (*User, Error) {
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("firstName, lastName, and birthdate are required")
	}
	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthdate,
		createdAt: time.Now(),
	}, nil
}
