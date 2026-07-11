package main

import (
	"fmt"

	"example.com/structs/admin"
	"example.com/structs/user"
)

func main() {

	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	// ... do something awesome with that gathered data!
	aUser := user.User{
		FirstName: firstName,
		LastName:  lastName,
		BirthDate: birthdate,
		//createdAt: time.Now(),
	}

	admin := admin.New("foo@example.com", "password")

	// this Email property is exposed because it's capital
	// So STRANGE
	fmt.Println(admin.Email)

	//fmt.Println(firstName, user.lastName, birthdate)
	aUser.SpeakMyName()
	aUser.ClearMyName()
	aUser.SpeakMyName()

	jack, error := user.New("Jack", "Sargent", "12/20/2021")

	if error != nil {
		panic("this didn't go well.")
	}

	fmt.Println(jack.FirstName)

}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
