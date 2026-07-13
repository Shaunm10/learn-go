package main

import (
	"fmt"

	"example.com/note-taking-exercise/note"
)

func main() {
	note := promptForNote()
	saveNoteToDisk(note)
}

func promptForNote() note.Note {
	title := ""
	content := ""
	fmt.Print("Note Title:")
	fmt.Scan(&title)
	fmt.Print("Note Content:")
	fmt.Scan(&content)

	note := note.New(title, content)

	return *note
}

func saveNoteToDisk(note note.Note) {
	note.SaveNote()
}
