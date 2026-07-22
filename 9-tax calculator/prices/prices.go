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
	TaxIncludedPrices map[string]string
}

func (job *TaxIncludedPriceJob) Process() {
	result := make(map[string]string)
	job.loadData()

	for _, price := range job.InputPrices {
		// perform calculation
		taxIncludedPrice := price * (1 + job.TaxRate)

		// set the value formatted to 2 decimal points
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}

	job.TaxIncludedPrices = result
	filePathForJob := fmt.Sprintf("./output/result_%.0f.json", job.TaxRate*100)
	err := fileManager.WriteJSON(job, filePathForJob)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("File successfully exported: " + filePathForJob)
	}
}

func (job *TaxIncludedPriceJob) loadData() {

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
	job.InputPrices = append(job.InputPrices, prices...)
}

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
	//defaultInputPrices := []float64{10, 20, 30}

	return &TaxIncludedPriceJob{
		TaxRate: taxRate,
		//InputPrices: defaultInputPrices,
	}
}
