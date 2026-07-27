package models

import (
	"errors"
	"time"

	"example.com/rest-api/database"
	"example.com/rest-api/utils"
)

type User struct {
	ID          int64
	Email       string `binding:"required"`
	Password    string `binding:"required"`
	CreatedDate time.Time
}

// TODO: move this to a services library
func (user *User) Save() error {
	query := `
	INSERT INTO users(email, password, createDate)
	VALUES (?,?,datetime('now'))
	`
	stmt, err := database.DatabaseInstance.Prepare(query)

	if err != nil {
		return errors.Join(errors.New("Unable to prepare events INSERT statement"), err)
	}

	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(user.Password)

	if err != nil {
		return errors.Join(errors.New("Unable to calculate password hash"), err)
	}

	result, err := stmt.Exec(user.Email, hashedPassword)

	if err != nil {
		return errors.Join(errors.New("Failed to INSERT event"), err)
	}

	userId, err := result.LastInsertId()

	if err != nil {
		return errors.Join(errors.New("Unable to fetch newly created UserId"), err)
	}
	user.ID = userId

	return nil
}

func GetHashedEmailByEmail(email string) (string, error) {
	query := `
		SELECT 
			 Password
		FROM users
		WHERE Email = ?
	`

	//var user User
	var retrievedPassword string
	row := database.DatabaseInstance.QueryRow(query, email)

	err := row.Scan(&retrievedPassword)

	if err != nil {
		return retrievedPassword, err
	}

	return retrievedPassword, err
}
