package prices

import "fmt"

type TaxIncludedPriceJob struct {
	TaxRate     float64
	InputPrices []float64
	// The key will be the input price
	TaxIncludedPrices map[string]float64
}

func (taxIncludedPriceJob *TaxIncludedPriceJob) Process() {
	result := make(map[string]float64)

	for _, price := range taxIncludedPriceJob.InputPrices {
		result[fmt.Sprintf("%.2f", price)] = price * (1 + taxIncludedPriceJob.TaxRate)
	}
	fmt.Println(result)
	//taxIncludedPriceJob.TaxIncludedPrices = result
}

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
	defaultInputPrices := []float64{10, 20, 30}

	return &TaxIncludedPriceJob{
		TaxRate:     taxRate,
		InputPrices: defaultInputPrices,
	}
}
