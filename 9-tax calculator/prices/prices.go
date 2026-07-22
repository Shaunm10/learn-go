package prices

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

	var linesBuffer []string
	for scanner.Scan() {
		linesBuffer = append(linesBuffer, scanner.Text())
	}

	// did the scanner have an error?
	err = scanner.Err()
	if err != nil {
		// TODO: handle gracefully
		fmt.Println("Error scanning file", err)
		file.Close()
		return
	}
	file.Close()

	prices := make([]float64, len(linesBuffer))
	// iterate thought the strings converting them.
	for lineIndex, line := range linesBuffer {
		floatPrice, err := strconv.ParseFloat(line, 64)
		if err != nil {
			// if one price conversion failed, stop processing
			fmt.Printf("Unable to convert value '%v' to a valid Float64 \n\n", line)
			fmt.Println(err)
			file.Close()

			return
		} else {

			prices[lineIndex] = floatPrice
			//prices = append(prices, floatPrice)
		}
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
