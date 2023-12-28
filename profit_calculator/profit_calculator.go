package main

import (
	"fmt"
)

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64
	gettingInput("Revenue", &revenue)
	gettingInput("Expenses", &expenses)
	gettingInput("Tax Rate", &taxRate)

	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit

	fmt.Println(ebt)
	fmt.Println(profit)
	fmt.Println(ratio)
}

func gettingInput(input string, result *float64) {
	fmt.Printf("%s: ", input)
	fmt.Scan(result)
}
