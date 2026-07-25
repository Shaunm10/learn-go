package models

import (
	"database/sql"
	"errors"
	"time"

	"example.com/rest-api/database"
)

type Event struct {
	ID          int64
	Name        string    `binding:required`
	Description string    `binding:required`
	Location    string    `binding:required`
	DateTime    time.Time `binding:required`
	UserID      int
}

func (event *Event) Save() error {

	query := `
	INSERT INTO events(name, description, location, dateTime, user_Id)
	VALUES (?,?,?,?,?)
	`
	stmt, err := database.DatabaseInstance.Prepare(query)

	if err != nil {
		return errors.Join(errors.New("Unable to prepare events INSERT statement"), err)
	}

	defer stmt.Close()

	result, err := stmt.Exec(event.Name, event.Description, event.Location, event.DateTime, event.UserID)

	if err != nil {
		return errors.Join(errors.New("Failed to INSERT event"), err)
	}

	eventId, err := result.LastInsertId()

	if err != nil {
		return errors.Join(errors.New("Unable to fetch newly created EventId"), err)
	}
	event.ID = eventId

	return nil
}

func GetAllEvents() ([]Event, error) {
	query := `
		SELECT
			id
			, name
			, description
			, location
			, dateTime
			, user_Id
		FROM events
	`

	rows, err := database.DatabaseInstance.Query(query)

	if err != nil {
		return nil, errors.Join(errors.New("Unable to prepare events SELECT statement"), err)
	}

	defer rows.Close()
	var returnEvents []Event

	for rows.Next() {
		var currentEvent Event
		err := rows.Scan(&currentEvent.ID, &currentEvent.Name, &currentEvent.Description, &currentEvent.Location, &currentEvent.DateTime, &currentEvent.UserID)

		if err != nil {
			return nil, errors.Join(errors.New("Unable to populate event from data"), err)
		}
		returnEvents = append(returnEvents, currentEvent)
	}

	err = rows.Err()
	if err != nil {
		return nil, errors.Join(errors.New("rows reported an error, so unable to populate event from data"), err)
	}

	return returnEvents, nil
}

func GetEvent(eventId int64) (*Event, error) {
	query := `
		SELECT
			id
			, name
			, description
			, location
			, dateTime
			, user_Id
		FROM events
		WHERE id =?
	`
	row := database.DatabaseInstance.QueryRow(query, eventId)

	var event Event

	err := row.Scan(
		&event.ID,
		&event.Name,
		&event.Description,
		&event.Location,
		&event.DateTime,
		&event.UserID)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// we couldn't find the 1 row so this is a http 404
			return nil, nil
		}

		// an actual DB/SQL error
		return nil, errors.Join(
			errors.New("Unable to get event for Id"),
			err,
		)
	}

	// happy path
	return &event, nil
}
