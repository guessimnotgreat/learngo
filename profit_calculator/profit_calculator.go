package main

import (
	"fmt"
)

func main() {
	revenue := gettingInput("Revenue")
	expenses := gettingInput("Expenses")
	taxRate := gettingInput("Tax Rate")

	ebt, profit, ratio := calculateFinancial(revenue, expenses, taxRate)

	fmt.Printf("%.1f\n", ebt)
	fmt.Printf("%.1f\n", profit)
	fmt.Printf("%.3f\n", ratio)
}

func gettingInput(input string) float64 {
	var result float64
	fmt.Printf("%s: ", input)
	fmt.Scan(&result)
	return result
}

func calculateFinancial(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}
