package main

import (
	"example.com/price-calculator/prices"
	terminalmanager "example.com/price-calculator/terminalManager"
)

func main() {

	taxRates := []float64{0, 0.07, 0.1, 0.15}

	for _, taxRate := range taxRates {
		//outputFilePath := fmt.Sprintf("./output/result_%.0f.json", taxRate*100)
		//fileManager := fileManager.New("prices.txt", outputFilePath)
		termManager := terminalmanager.New()
		priceJob := prices.NewTaxIncludedPriceJob(termManager, taxRate)
		//priceJob := prices.NewTaxIncludedPriceJob(fileManager, taxRate)

		priceJob.Process()
	}
}
