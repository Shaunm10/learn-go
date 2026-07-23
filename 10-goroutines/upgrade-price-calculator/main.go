package main

import (
	"fmt"

	"example.com/price-calculator/filemanager"
	"example.com/price-calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	// create an "slice" of channels with the size of our loop
	doneChans := make([]chan bool, len(taxRates))
	errorChans := make([]chan error, len(taxRates))
	for index, taxRate := range taxRates {

		// for each iteration create a new channel and set it to the index
		doneChans[index] = make(chan bool)
		errorChans[index] = make(chan error)
		fm := filemanager.New("prices.txt", fmt.Sprintf("./output/result_%.0f.json", taxRate*100))
		// cmdm := cmdmanager.New()
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)

		go priceJob.Process(doneChans[index], errorChans[index])

		// if err != nil {
		// 	fmt.Println("Could not process job")
		// 	fmt.Println(err)
		// }
	}
	// go through all the iterations the loop did
	for index := range taxRates {
		select {
		// once one of these cases executes, it won't wait for the other (done vs error)
		case err := <-errorChans[index]:
			if err != nil {
				// do something with the error
				fmt.Println("Could not process job")
				fmt.Println(err)
			}

		case <-doneChans[index]:
			// do something when each item is done
			fmt.Println("Done!")
		}
	}

	fmt.Println("Complete!")
}
