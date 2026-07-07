package main

import "fmt"

/*
	GOALS:
	- refactor the prompt for input into different functions.
	- refactor the calculation of variables into a different function.
*/

func main() {
	//var revenue float64
	//var expenses float64
	//var taxRate float64

	revenue := promptForFloatInput("Revenue:")
	// fmt.Print("Revenue: ")
	// fmt.Scan(&revenue)

	expenses := promptForFloatInput("Expenses:")
	// fmt.Print("Expenses: ")
	// fmt.Scan(&expenses)

	// fmt.Print("Tax Rate: ")
	// fmt.Scan(&taxRate)
	taxRate := promptForFloatInput("Tax Rate:")

	ebt, profit, ratio := performCalculations(revenue, expenses, taxRate)
	// ebt := revenue - expenses
	// profit := ebt * (1 - taxRate/100)
	// ratio := ebt / profit

	// fmt.Println(ebt)
	// fmt.Println(profit)
	// fmt.Println(ratio)
	fmt.Printf(`
Financial Results:
	EBT: $%.2f
	Profit: $%.2f
	Ratio: %.6f
	`, ebt, profit, ratio)
}

func promptForFloatInput(prompt string) (capturedValue float64) {
	fmt.Print(prompt)
	fmt.Scan(&capturedValue)
	return
}

func performCalculations(revenue, expenses, taxRate float64) (ebt, profit, ratio float64) {
	ebt = revenue - expenses
	profit = ebt * (1 - taxRate/100)
	ratio = ebt / profit
	return
}
