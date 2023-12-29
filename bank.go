package main

import (
	"fmt"

	"github.com/guessimnotgreat/learngo/fileops"
)

const accountBalanceFile = "balance.txt"

func main() {

	accountBalance, err := fileops.GetFloatFromFile(accountBalanceFile)
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("----------")
		panic("Can't continue, sorry.")
	}

	printMenu()

	for {
		choice := getUserChoice()

		switch choice {
		case 1:
			fmt.Println("Your balance is", accountBalance)
		case 2:
			var depositAmount float64
			fmt.Print("Your deposit: ")
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0")
				continue
			}

			accountBalance += depositAmount
			fmt.Println("Balance updated! New amount:", accountBalance)
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
		case 3:
			var withdrawAmount float64
			fmt.Print("Your withdraw: ")
			fmt.Scan(&withdrawAmount)

			if withdrawAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0")
				continue
			}

			if withdrawAmount > accountBalance {
				fmt.Println("Invalid amount. You can't withdraw more than you have.")
				continue
			}

			accountBalance -= withdrawAmount
			fmt.Println("Balance updated! New amount:", accountBalance)
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
		default:
			fmt.Println("Goodbye!")
			return
		}
	}

}

func getUserChoice() int {
	var choice int
	fmt.Print("Your choice: ")
	fmt.Scan(&choice)
	return choice
}
