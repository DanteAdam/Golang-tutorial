package prices

import "fmt"

type taxIncludedPricesJob struct {
	TaxRate           float64
	InputPrices       []float64
	taxIncludedPrices map[string]float64
}

func (job taxIncludedPricesJob) Process() {
	result := make(map[string]float64)

	for _, price := range job.InputPrices {
		result[fmt.Sprintf("%.2f", price)] = price * (1 + job.TaxRate)
	}
	fmt.Println(result)
}

func NewTaxIncludedPriceJob(taxRate float64) *taxIncludedPricesJob {
	return &taxIncludedPricesJob{
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}
