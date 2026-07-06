package main

import (
	"fmt"
	"math"
	"strconv"
)

/*
Expected Output:

Investment Amount: 10000
Expected Return Rate: 5.5
Years: 10
Future Value -> 17081.4
Future Value -> (adjusted for Inflation): 10446.2
*/
func main() {

	const inflationRate = 2.5
	// this is an example, however not very readable, so this should be split up.
	var investmentAmount, years, expectedReturnRate = 1000.0, 10.0, 5.5
	////expectedReturnRate := 5.5
	//years := 10

	// prompt for Investment Amount
	fmt.Print("What is the investment amount:")
	// populates the variable (via pointer) to the inputted value
	fmt.Scan(&investmentAmount)

	// prompt for number of years
	fmt.Print("How many years is the investment for:")
	fmt.Scan(&years)

	// prompt for expected return rate
	fmt.Print("What is the expected return rate:")
	fmt.Scan(&expectedReturnRate)

	var percentageIncrease = (1 + expectedReturnRate/100)
	var futureValue = investmentAmount * math.Pow((percentageIncrease), years)
	var futureValueAdjustedForInflation = futureValue / math.Pow((1+(inflationRate/100)), years)

	fmt.Println("The future value of the investment is: " + strconv.FormatFloat(futureValue, 'f', 2, 64))
	fmt.Println("The future value of the investment factoring inflation is: " + strconv.FormatFloat(futureValueAdjustedForInflation, 'f', 2, 64))

}
