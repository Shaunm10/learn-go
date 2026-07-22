package conversion

import (
	"errors"
	"fmt"
	"strconv"
)

// Converts a slice of strings to a slice of of floats. It can also return an error if
// conversion fails
func StringsToFloats(strings []string) ([]float64, error) {
	floatsToReturn := make([]float64, len(strings))

	for stringIndex, stringValue := range strings {

		floatConverted, err := strconv.ParseFloat(stringValue, 64)
		if err != nil {

			errorToThrow := errors.New(fmt.Sprintf("Unable to convert value %s to a float64", stringValue))
			return nil, errorToThrow
		}

		floatsToReturn[stringIndex] = floatConverted
	}

	return floatsToReturn, nil
}
