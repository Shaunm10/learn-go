package models

import "time"

type Event struct {
	ID          int
	Name        string    `binding:required`
	Description string    `binding:required`
	Location    string    `binding:required`
	DateTime    time.Time `binding:required`
	UserID      int
}

func (event Event) Save() {
	events = append(events, event)
}

// static fields
var events = []Event{}

func GetAllEvents() []Event {
	return events
}
