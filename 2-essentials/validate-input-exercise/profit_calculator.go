package main

import (
	"errors"
	"fmt"
	"os"
)

/*
Goals
--------

	1) Validate user input
	=> show error message & exit if invalid input is provided
	- No negative numbers
	- Not 0

	2) store calculated results into file

*/

func main() {
	revenue, err := getUserInput("Revenue: ")
	if err != nil {
		showCatastrophicError("Revenue", err)
	}

	expenses, err := getUserInput("Expenses: ")
	if err != nil {
		showCatastrophicError("Expenses", err)
	}

	taxRate, err := getUserInput("Tax Rate: ")
	if err != nil {
		showCatastrophicError("Tax Rate", err)
	}

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	fmt.Printf("%.1f\n", ebt)
	fmt.Printf("%.1f\n", profit)
	fmt.Printf("%.3f\n", ratio)

	saveData(ebt, profit, ratio)
}

func saveData(ebt, profit, ratio float64) {
	var dataToPersist = fmt.Sprintf("Ebt: %.2f, Profit: %.2f, ratio: %.2f", ebt, profit, ratio)
	os.WriteFile("financialResults.data", []byte(dataToPersist), 0644)
}

func showCatastrophicError(propertyName string, err error) {
	fmt.Printf("Error getting %s: %s\n\n", propertyName, err.Error())
	panic("Unrecoverable error, shutting down.")
}

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}

func getUserInput(infoText string) (float64, error) {
	var userInput float64
	fmt.Print(infoText)
	_, err := fmt.Scan(&userInput)

	if err != nil {
		errorMessage := fmt.Sprintf("Unacceptable value, `%s` is not convertible to a numeric value.")
		return 0.0, errors.New(errorMessage)
	}

	if userInput <= 0 {
		errorMessage := fmt.Sprintf("Unacceptable value, it must be greater than 0 and you entered '%v'", userInput)
		return 0.0, errors.New(errorMessage)
	}
	return userInput, nil
}
