package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	FirstName string
	LastName  string
	BirthDate string
	createdAt time.Time
}

func (user User) SpeakMyName() {

	fmt.Println(user.FirstName, user.LastName, user.BirthDate)
}

func (user *User) ClearMyName() {
	user.FirstName = ""
	user.LastName = ""
}

func New(firstName, lastName, birthdate string) (*User, error) {
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("firstName, lastName, and birthdate are required")
	}
	return &User{
		FirstName: firstName,
		LastName:  lastName,
		BirthDate: birthdate,
		createdAt: time.Now(),
	}, nil
}
