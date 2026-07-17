package note

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title        string    `json:"title"`
	Content      string    `json:"content"`
	CreationDate time.Time `json:"creationDate"`
}

func New(title, content string) *Note {

	return &Note{
		Title:        title,
		Content:      content,
		CreationDate: time.Now(),
	}
}

func (note Note) SaveNote() error {
	subDirectory := "savedNotes/"
	lowerFileName := strings.ToLower(note.Title)
	fileName := subDirectory + strings.ReplaceAll(lowerFileName, " ", "_") + ".json"
	bytes, err := json.Marshal(note)

	if err != nil {
		return err

	}

	return os.WriteFile(fileName, bytes, 0644)
}

func (note Note) Display() {
	fmt.Println()
	fmt.Println("**********************************")
	fmt.Println("Title: ", note.Title)
	fmt.Println("Content: ", note.Content)
	fmt.Println("**********************************")

	//fmt.Printf("Your note title %v has the following content:\n\n", note.Title, note.Content)

}
