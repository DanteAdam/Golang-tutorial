package main

import (
	"fmt"

	"example.com/price-cal/filemanager"
	"example.com/price-cal/prices"
)

type IOManager interface {
	ReadLines() ([]string, error)
	WriteResult(data any) error
}

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}
	doneChans := make([]chan bool, len(taxRates))
	errorChans := make([]chan error, len(taxRates))

	for index, taxRate := range taxRates {
		doneChans[index] = make(chan bool)
		errorChans[index] = make(chan error)

		fm := filemanager.NewFileManager("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		// cmd := cmdmanager.NewCMDManager()
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		go priceJob.Process(doneChans[index], errorChans[index])

		// if err != nil {
		// 	fmt.Println("Could not process job")
		// 	fmt.Println(err)
		// }
	}
	for index, _ := range taxRates {
		select {
		case err := <-errorChans[index]:
			if err != nil {
				fmt.Println(err)
			}
		case <-doneChans[index]:
			fmt.Println("Done")
		}

	}

}
