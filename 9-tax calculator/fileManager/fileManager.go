package fileManager

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type FileManager struct {
	InputFilePath  string
	OutputFilePath string
}

func (fm FileManager) ReadLines() ([]string, error) {

	file, err := os.Open(fm.InputFilePath)

	if err != nil {

		//fmt.Println("Error opening price.txt", err)
		// TODO: handle gracefully
		return nil, errors.New("Failed to open file.")
	}

	// otherwise try to readin the file contents
	scanner := bufio.NewScanner(file)

	// close the file whenever this method goes out of scope.
	defer file.Close()

	var linesBuffer []string
	for scanner.Scan() {
		linesBuffer = append(linesBuffer, scanner.Text())
	}

	// did the scanner have an error?
	err = scanner.Err()
	if err != nil {
		return nil, fmt.Errorf("Scanner returned error: %w", err)
	}

	return linesBuffer, nil
}

// writes an object to a specific path
func (fm FileManager) WriteResult(data any) error {

	if fm.OutputFilePath == "" {
		return errors.New("filePath must be specified")
	}

	file, err := os.Create(fm.OutputFilePath)

	if err != nil {
		return errors.Join(errors.New("unable to create a file"), err)
	}

	// close this file when this method exits
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)

	if err != nil {
		return errors.Join(errors.New("unable to write Json to file"), err)
	}

	// happy path
	return nil
}

func New(inputFilePath, outputFilePath string) FileManager {
	// validation??
	return FileManager{
		InputFilePath:  inputFilePath,
		OutputFilePath: outputFilePath,
	}
}
