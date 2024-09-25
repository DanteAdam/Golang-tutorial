package cmdmanager

import "fmt"

type CMDManager struct {
}

func (cmd CMDManager) ReadLines() ([]string, error) {
	fmt.Println("Enter your prices. Confirm every price with ENTER")

	prices := []string{}

	for {
		price := ""
		fmt.Print("Price:")
		fmt.Scan(&price)

		if price == "0" {
			break
		}
		prices = append(prices, price)
	}
	return prices, nil
}

func (cmd CMDManager) WriteResult(data any) error {
	fmt.Println(data)
	return nil
}

func NewCMDManager() CMDManager {
	return CMDManager{}
}
