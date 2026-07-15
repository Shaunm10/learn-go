package main

import (
	"bufio"
	"fmt"
	"os"

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
	title = inputSentence()
	//fmt.Scanln(&title)
	fmt.Print("Note Content:")
	content = inputSentence()
	//fmt.Scanln(&content)

	note := note.New(title, content)

	return *note
}

func saveNoteToDisk(note note.Note) {
	note.SaveNote()
}

func inputSentence() string {
	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		sentence := scanner.Text()
		return sentence
	}
	return ""
}
