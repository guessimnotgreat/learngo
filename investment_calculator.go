package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 6.5
	var investmentAmount float64
	var years float64
	expectedReturnRate := 5.5

	gettingInput("Investment Amount", &investmentAmount)
	gettingInput("Expected Return Rate", &expectedReturnRate)
	gettingInput("Years", &years)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}

func gettingInput(input string, result *float64) {
	fmt.Printf("%s: ", input)
	fmt.Scan(result)
}
