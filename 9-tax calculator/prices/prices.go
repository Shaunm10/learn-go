package prices

import (
	"fmt"

	"example.com/price-calculator/conversion"
	"example.com/price-calculator/fileManager"
)

const pricesFilePath = "prices.txt"

type TaxIncludedPriceJob struct {
	TaxRate     float64
	InputPrices []float64
	// The key will be the input price
	TaxIncludedPrices map[string]float64
}

func (taxIncludedPriceJob *TaxIncludedPriceJob) Process() {
	result := make(map[string]string)
	taxIncludedPriceJob.loadData()

	for _, price := range taxIncludedPriceJob.InputPrices {
		// perform calculation
		taxIncludedPrice := price * (1 + taxIncludedPriceJob.TaxRate)

		// set the value formatted to 2 decimal points
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}
	fmt.Println(result)
	//taxIncludedPriceJob.TaxIncludedPrices = result
}

func (taxIncludedPriceJob *TaxIncludedPriceJob) loadData() {

	linesFromFile, err := fileManager.ReadLines(pricesFilePath)

	if err != nil {
		fmt.Println("Error scanning file", err)
		return
	}

	prices, err := conversion.StringsToFloats(linesFromFile)

	if err != nil {
		fmt.Println(err)
		return
	}

	// finally populate the input prices
	taxIncludedPriceJob.InputPrices = append(taxIncludedPriceJob.InputPrices, prices...)
}

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
	//defaultInputPrices := []float64{10, 20, 30}

	return &TaxIncludedPriceJob{
		TaxRate: taxRate,
		//InputPrices: defaultInputPrices,
	}
}
