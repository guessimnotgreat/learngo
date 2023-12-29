package main

import (
	"fmt"
	"os"
	"strconv"
)

const (
	accountBalanceFile = "balance.txt"
	WelcomeMessage     = "Welcome to Go Bank!"
	PromptMessage      = "What do you want to do?"
	CheckBalance       = "1. Check balance"
	DepositMoney       = "2. Deposit money"
	WithdrawMoney      = "3. Withdraw money"
	ExitMessage        = "4. Exit"
)

func main() {

	accountBalance := getBalanceFromFile()

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
			writeBalanceToFile(accountBalance)
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
			writeBalanceToFile(accountBalance)
		default:
			fmt.Println("Goodbye!")
			return
		}
	}

}

func printMenu() {
	items := []string{
		WelcomeMessage,
		PromptMessage,
		CheckBalance,
		DepositMoney,
		WithdrawMoney,
		ExitMessage,
	}

	for _, item := range items {
		fmt.Println(item)
	}
}

func getUserChoice() int {
	var choice int
	fmt.Print("Your choice: ")
	fmt.Scan(&choice)
	return choice
}

func writeBalanceToFile(balance float64) {
	balanceTxt := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceTxt), 0644)
}

func getBalanceFromFile() float64 {
	balanceTxt, _ := os.ReadFile(accountBalanceFile)
	balance, _ := strconv.ParseFloat(string(balanceTxt), 64)
	return balance
}
