package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func ReadFloatFromFile(fileName string, defaultValue float64) (float64, error) {
	data, err := os.ReadFile(fileName)

	if err != nil {
		return defaultValue, errors.New("Failed to find file.")
	}

	valueText := string(data)
	value, err := strconv.ParseFloat(valueText, 64)

	if err != nil {
		return defaultValue, errors.New("Failed to parse stored value.")
	}

	return value, nil
}

func WriteFloatToFile(value float64, filePath string) {
	valueText := fmt.Sprint(value)
	os.WriteFile(filePath, []byte(valueText), 0644)
}
