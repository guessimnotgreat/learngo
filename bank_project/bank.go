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
	printMenu(
		WelcomeMessage,
		PromptMessage,
		CheckBalance,
		DepositMoney,
		WithdrawMoney,
		ExitMessage,
	)

	var choice int
	fmt.Print("Your choice: ")
	fmt.Scan(&choice)

	fmt.Println("Your choice:", choice)
}

func printMenu(items ...string) {
	for _, item := range items {
		fmt.Println(item)
	}
}
