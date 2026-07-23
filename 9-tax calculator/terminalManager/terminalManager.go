package terminalmanager

import (
	"fmt"
	"strings"
)

type TerminalManager struct {
}

func (terminalManager TerminalManager) ReadLines() ([]string, error) {
	input := ""

	fmt.Println("Please enter in some prices separated by a comma")
	fmt.Scan(&input)
	lines := strings.Split(input, ",")
	return lines, nil
}

func (terminalManager TerminalManager) WriteResult(data any) error {
	fmt.Println(data)
	return nil
}

func New() TerminalManager {
	return TerminalManager{}
}
