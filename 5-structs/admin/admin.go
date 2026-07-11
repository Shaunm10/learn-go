package admin

import (
	"example.com/structs/user"
)

type Admin struct {
	Email    string
	password string
	User     user.User
}

func New(email, password string) Admin {

	return Admin{
		Email:    email,
		password: password,
		User: user.User{
			FirstName: "ADMIN",
			LastName:  "ADMIN",
			BirthDate: "---",
		},
	}
}
