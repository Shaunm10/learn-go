package note

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title        string
	Content      string
	creationDate time.Time
}

func New(title, content string) *Note {

	return &Note{
		Title:        title,
		Content:      content,
		creationDate: time.Now(),
	}
}

func (note Note) SaveNote() {
	lowerFileName := strings.ToLower(note.Title)
	fileName := strings.ReplaceAll(lowerFileName, " ", "_") + ".json"
	bytes, err := json.Marshal(note)

	if err != nil {
		fmt.Println("Unable to save file.")
		return
	}

	os.WriteFile(fileName, bytes, 0644)

}
