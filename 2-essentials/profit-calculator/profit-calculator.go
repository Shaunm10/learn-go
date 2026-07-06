package main

import (
	"fmt"
	"strconv"
)

func main() {

	revenue, expenses, taxRate := 0.0, 0.0, 0.0

	fmt.Println("===============================================")
	fmt.Println("Welcome to the profit calculator")
	fmt.Println("===============================================")

	fmt.Println("")
	fmt.Println("")

	// prompts
	// revenue
	fmt.Println("What has your revenue been?:")
	fmt.Scan(&revenue)

	// expenses
	fmt.Println("What has your expenses been?:")
	fmt.Scan(&expenses)

	// tax rate
	fmt.Println("What is your tax rate?:")
	fmt.Scan(&taxRate)

	// calculations

	ebt := revenue - expenses
	//taxAmount := ebt * (taxRate / 100)
	profit := ebt * (1 - (taxRate / 100))
	earningsAfterTax := ebt - taxRate
	ratio := ebt / profit

	// output
	fmt.Println("Your Earnings before Tax: $" + strconv.FormatFloat(ebt, 'f', 2, 64))
	fmt.Println("Your Earnings after Tax: $" + strconv.FormatFloat(earningsAfterTax, 'f', 2, 64))
	fmt.Println("Your ratio of revenue to profit is: %" + strconv.FormatFloat(ratio, 'f', 2, 64))
}
