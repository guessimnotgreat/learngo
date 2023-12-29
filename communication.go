package main

import (
	"fmt"

	"github.com/Pallinder/go-randomdata"
)

const (
	WelcomeMessage = "Welcome to Go Bank!"
	PromptMessage  = "What do you want to do?"
	CheckBalance   = "1. Check balance"
	DepositMoney   = "2. Deposit money"
	WithdrawMoney  = "3. Withdraw money"
	ExitMessage    = "4. Exit"
)

func printMenu() {
	callUsMessage := fmt.Sprintf("Reach us 24/7 at %v", randomdata.PhoneNumber())
	items := []string{
		WelcomeMessage,
		callUsMessage,
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
