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
	for index, taxRate := range taxRates {

		// for each iteration create a new channel and set it to the index
		doneChans[index] = make(chan bool)
		fm := filemanager.New("prices.txt", fmt.Sprintf("./output/result_%.0f.json", taxRate*100))
		// cmdm := cmdmanager.New()
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)

		go priceJob.Process(doneChans[index])

		// if err != nil {
		// 	fmt.Println("Could not process job")
		// 	fmt.Println(err)
		// }
	}

	// now wait for each go routine to finish
	for _, doneChan := range doneChans {
		<-doneChan
	}

	fmt.Println("Complete!")
}
