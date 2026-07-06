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
	var investmentAmount = 1000
	var expectedReturnRate = 5.5
	var years = 10

	var percentageIncrease = (1 + expectedReturnRate/100)
	var futureValue = float64(investmentAmount) * math.Pow((percentageIncrease), float64(years))

	fmt.Println("The future value of the investment is: " + strconv.FormatFloat(futureValue, 'f', -1, 64))
}
