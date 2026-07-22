package prices

import (
	"bufio"
	"fmt"
	"os"

	"example.com/price-calculator/conversion"
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

	file, err := os.Open(pricesFilePath)

	if err != nil {
		fmt.Println("Error opening price.txt", err)
		// TODO: handle gracefully
		return
	}

	scanner := bufio.NewScanner(file)

	// close the file whenever.
	defer file.Close()

	var linesBuffer []string
	for scanner.Scan() {
		linesBuffer = append(linesBuffer, scanner.Text())
	}

	// did the scanner have an error?
	err = scanner.Err()
	if err != nil {
		// TODO: handle gracefully
		fmt.Println("Error scanning file", err)
		//file.Close()
		return
	}
	//file.Close()

	prices, err := conversion.StringsToFloats(linesBuffer)

	if err != nil {
		fmt.Println(err)
		//file.Close()
		return
	}

	//file.Close()

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
