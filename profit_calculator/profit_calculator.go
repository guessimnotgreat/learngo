package main

import (
	"fmt"
	"os"
)

const resultFileName = "result.txt"

func main() {
	revenue := gettingInput("Revenue")
	expenses := gettingInput("Expenses")
	taxRate := gettingInput("Tax Rate")

	ebt, profit, ratio := calculateFinancial(revenue, expenses, taxRate)

	writeToFileAndDisplay(ebt, profit, ratio)
}

func gettingInput(input string) float64 {
	var result float64
	fmt.Printf("%s: ", input)
	fmt.Scan(&result)

	if result <= 0 {
		errMessage := fmt.Sprintf("%v must not be negative or 0", input)
		panic(errMessage)
	}

	return result
}

func calculateFinancial(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}

func writeToFileAndDisplay(ebt, profit, ratio float64) {
	resultTxt := fmt.Sprintf("EBT: %.1f\nPROFIT: %.1f\nRATIO: %.3f\n", ebt, profit, ratio)
	fmt.Print(resultTxt)
	os.WriteFile(resultFileName, []byte(resultTxt), 0644)
}
