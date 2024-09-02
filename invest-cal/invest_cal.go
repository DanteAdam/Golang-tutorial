package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 6.5
	var investmentAmount float64
	var years float64
	var expectedReturnRate float64

	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRevalue := futureValue / math.Pow(1+inflationRate/100, years)

	// outputs information
	// fmt.Println("Future Value:", futureValue)
	fmt.Printf("Future Value: %.2f\nFuture Value: %.2f\n", futureValue, futureRevalue)
	// fmt.Println("Future Revalue:", futureRevalue)

}
