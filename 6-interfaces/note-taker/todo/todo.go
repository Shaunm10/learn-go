package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
)

type Todo struct {
	//Title     string    `json:"title"`
	Text string `json:"text"`
	//CreatedAt time.Time `json:"created_at"`
}

func (todo Todo) Display() {
	//fmt.Printf("Your note titled %v has the following content:\n\n%v\n\n", todo.Title, todo.Text)
	fmt.Printf("Your todo is: \n\n%v\n\n", todo.Text)

}

func (todo Todo) Save() error {
	fileName := strings.ReplaceAll(todo.Text, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"

	json, err := json.Marshal(todo)

	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0644)
}

// Convince constructor
func New(text string) (Todo, error) {
	if text == "" {
		return Todo{}, errors.New("Invalid input.")
	}

	return Todo{
		Text: text,
	}, nil
}
