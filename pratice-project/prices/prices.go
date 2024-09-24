package prices

import (
	"fmt"

	"example.com/price-cal/conversion"
	"example.com/price-cal/filemanager"
)

type taxIncludedPricesJob struct {
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]string
}

func (job *taxIncludedPricesJob) LoadData() {

	lines, err := filemanager.ReadLines("prices.txt")

	if err != nil {
		fmt.Println(err)
		return
	}

	prices, err := conversion.StringsToFloats(lines)

	if err != nil {
		fmt.Println(err)
		return
	}
	job.InputPrices = prices
}

func (job *taxIncludedPricesJob) Process() {
	job.LoadData()

	result := make(map[string]string)

	for _, price := range job.InputPrices {
		taxIncludedPrices := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%2.f", taxIncludedPrices)
	}

	job.TaxIncludedPrices = result

	filemanager.WriteJSON(fmt.Sprintf("result_%.0f.json", job.TaxRate*100), job)
}

func NewTaxIncludedPriceJob(taxRate float64) *taxIncludedPricesJob {
	return &taxIncludedPricesJob{
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}
