package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"

	"example.com/note-taking-exercise/note"
)

func main() {
	note, error := promptForNote()
	if error != nil {
		panic(error)
	}
	note.Display()
	saveNoteToDisk(note)
}

func promptForNote() (note.Note, error) {
	title := ""
	content := ""
	fmt.Print("Note Title:")
	title = inputSentence()
	//fmt.Scanln(&title)
	fmt.Print("Note Content:")
	content = inputSentence()
	//fmt.Scanln(&content)

	if title == "" || content == "" {
		return note.Note{}, errors.New("title and content are required to create a Note.")
	}

	note := note.New(title, content)

	return *note, nil
}

func saveNoteToDisk(note note.Note) {
	err := note.SaveNote()
	if err != nil {
		panic(err)
	}

	fmt.Println("saving note successful")
}

func inputSentence() string {
	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		sentence := scanner.Text()
		return sentence
	}
	return ""
}
