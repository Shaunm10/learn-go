package main

import "fmt"

type Product struct {
	title string
	id    string
	price float64
}

func main() {
	// creates a fixed sized array of float64 with 4 values.
	prices := [4]float64{10.99, 9.99, 45.99, 20.0}

	// set the first item
	prices[0] = 99.9

	//thirdElement := prices[2]

	//lastTwoPrices := prices[2:4]

	// gets all the items up to the 3rd one.
	beginningPrices := prices[:3]

	// gets all the items index 2 and beyond
	lastPrices := prices[2:]

	numberOfItems := len(lastPrices)

	// dynamically sized array
	pricesOfAnySize := []float64{}
	// or initialized with values
	pricesOfAnySize2 := []float64{10.99, 79.99}

	// this will crash at runtime because this array doesn't yet have this index.
	pricesOfAnySize[2] = 10.99

	// append creates a new array but copying the existing AND adding 1 more item to the end.
	updatedPrices := append(pricesOfAnySize, 69.99)

	// removing the first element in the array is done via a slice
	pricesOfAnySize = pricesOfAnySize[1:]

	fmt.Println(updatedPrices)

}
