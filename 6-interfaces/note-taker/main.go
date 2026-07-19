package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
	"example.com/note/todo"
)

// in Go it's common to name the interface based on the method IF
// the interface only has 1 method, so {methodName} + "er"
// same rules apply for private vs public being camelCase or PascalCase
type saver interface {
	Save() error
}

type displayer interface {
	Display()
}

type outputtable interface {
	// embedded from another interface
	saver

	// new function definition
	Display()
}

func main() {

	title, content := getNoteData()

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	outputData(userNote)

	todoText := getUserInput("Todo Text: ")

	myTodo, err := todo.New(todoText)

	outputData(myTodo)
}

func saveData(note saver) error {
	err := note.Save()

	if err != nil {
		fmt.Println("Saving the note failed.")
		return err
	}
	fmt.Println("Saving the note succeeded!")
	return nil
}

func outputData(data outputtable) {

	err := data.Save()
	if err == nil {
		data.Display()
	}
}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")
	content := getUserInput("Note content:")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)

	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}

func printSomething(value interface{}) {

	// see if this value is an int
	typedVal, ok := value.(int)

	if ok {
		fmt.Println("This is an int: ", typedVal)
	}

	// NOTE: in these cases, the value's type doesn't get "narrowed" like in
	// TypeScript's type guards or use of typeof
	switch value.(type) {
	case int:
		fmt.Println("Integer:", value)
	case float64:
		fmt.Println("Float:", value)
	case string:
		fmt.Println("string:", value)
	default:
		// so some fallback logic

	}
}

// 'T' is the generic type, and since it's 'any' really
// any type can be passed in, but they must match. This will not compile
// because not every type can support the addition/ concatenation operator.
// func addAny[T any](a, b T) T {

// 	// won't compile
// 	return a + b
// }

// only int's float64 and strings can be passed,
// however the types for the parameters have to match
func add[T int | float64 | string](a, b T) T {
	return a + b
}
