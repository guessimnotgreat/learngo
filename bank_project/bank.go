package main

import "fmt"

const (
	WelcomeMessage = "Welcome to Go Bank!"
	PromptMessage  = "What do you want to do?"
	CheckBalance   = "1. Check balance"
	DepositMoney   = "2. Deposit money"
	WithdrawMoney  = "3. Withdraw money"
	ExitMessage    = "4. Exit"
)

func main() {

	accountBalance := 1000.00

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
				return
			}

			accountBalance += depositAmount
			fmt.Println("Balance updated! New amount:", accountBalance)
		case 3:
			var withdrawAmount float64
			fmt.Print("Your withdraw: ")
			fmt.Scan(&withdrawAmount)

			if withdrawAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0")
				return
			}

			if withdrawAmount > accountBalance {
				fmt.Println("Invalid amount. You can't withdraw more than you have.")
				return
			}

			accountBalance -= withdrawAmount
			fmt.Println("Balance updated! New amount:", accountBalance)
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
