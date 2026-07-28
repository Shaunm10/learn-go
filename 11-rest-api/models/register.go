package models

import (
	"errors"

	"example.com/rest-api/database"
)

type Register struct {
	ID      int64
	UserID  int64
	EventID int64
}

func CreateRegister(userId, eventId int64) (*Register, error) {

	query := `
		INSERT INTO registrations(user_id, event_id) 
		VALUES(?,?)
	`

	stmt, err := database.DatabaseInstance.Prepare(query)

	if err != nil {
		return nil, errors.Join(errors.New("Unable to prepare register INSERT statement"), err)
	}

	defer stmt.Close()

	result, err := stmt.Exec(userId, eventId)

	if err != nil {
		return nil, errors.Join(errors.New("Failed to INSERT register"), err)
	}

	registerId, err := result.LastInsertId()

	if err != nil {
		return nil, errors.Join(errors.New("Unable to fetch newly created registerId"), err)
	}

	newlyCreateRegister := Register{
		ID:      registerId,
		UserID:  userId,
		EventID: eventId,
	}

	return &newlyCreateRegister, nil
}

func DeleteRegister(userId, eventId int64) error {
	query := `
		DELETE 
		FROM registrations
		WHERE user_id = ?
		AND event_id = ?
	`

	_, err := database.DatabaseInstance.Exec(query, userId, eventId)

	if err != nil {
		return errors.Join(errors.New("Unable to DELETE registers"), err)
	}

	// happy path
	return nil
}
