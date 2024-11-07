package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func main() {
	var accountBalance, err = getBalanceToFile()

	if err != nil {
		fmt.Println("-------")
		fmt.Println("ERROR")
		fmt.Println(err)
		// Function for tracking the error infos
		// panic("Can't continue, sorry.")
		fmt.Println("-------")
	}

	for {
		initialQuestions()

		var choice int = getUsersChoice()

		switch choice {
		case 1:
			fmt.Println("Your balance is: ", accountBalance)
		case 2:
			accountBalance = depositAmount(accountBalance)
		case 3:
			accountBalance = withdrawAmount(accountBalance)
		case 4:
			fmt.Println("Goodbye!")
			fmt.Println("Thanks for choosing our bank.")
			return
		default:
			fmt.Println("Invalid choice.")
		}
		continue
	}
}

func initialQuestions() {
	fmt.Println("Welcome to Go Bank!")
	fmt.Println("What do you want to do?")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")
}

func getUsersChoice() int {
	fmt.Print("Your choice: ")
	var choice int
	fmt.Scan(&choice)
	return choice
}

func withdrawAmount(accountBalance float64) float64 {
	fmt.Print("Withdraw amount: ")
	var withdrawAmount float64
	fmt.Scan(&withdrawAmount)

	if withdrawAmount <= 0 {
		fmt.Println("Invalid amount. Must be greater than 0.")
	} else if withdrawAmount > accountBalance {
		fmt.Println("Invalid amount. You can't withdraw more than you have.")
	} else {
		accountBalance -= withdrawAmount
		fmt.Println("Balance updated! New amount:", accountBalance)
		writeBalanceToFile(accountBalance)
	}
	return accountBalance
}

func depositAmount(accountBalance float64) float64 {
	fmt.Print("Your deposit: ")
	var depositAmount float64
	fmt.Scan(&depositAmount)

	if depositAmount <= 0 {
		fmt.Println("Invalid amount. Must be greater than 0.")
	} else {
		accountBalance += depositAmount
		fmt.Println("Balance updated! New amount:", accountBalance)
		writeBalanceToFile(accountBalance)
	}
	return accountBalance
}

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
}

func getBalanceToFile() (float64, error) {
	// _ represents a value that I know I could receive but I don't want to work with it
	data, err := os.ReadFile(accountBalanceFile)
	if err != nil {
		return 0, errors.New("Failed to read the file.")
	}

	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)
	if err != nil {
		return 0, errors.New("Failed to parse stored balance value.")
	}

	return balance, nil
}
