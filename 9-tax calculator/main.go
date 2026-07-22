package main

import (
	"fmt"

	"example.com/price-calculator/fileManager"
	"example.com/price-calculator/prices"
)

func main() {

	taxRates := []float64{0, 0.07, 0.1, 0.15}

	for _, taxRate := range taxRates {
		filePathForJob := fmt.Sprintf("./output/result_%.0f.json", taxRate*100)
		fileManager := fileManager.New("prices.txt", filePathForJob)
		priceJob := prices.NewTaxIncludedPriceJob(fileManager, taxRate)

		priceJob.Process()
	}
}
