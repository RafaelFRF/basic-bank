package main

import (
	"fmt"

	"examples.com/bank/fileops"
)

const accountBalanceFile = "balance.txt"

func main() {
	var accountBalance, err = fileops.GetFloatFromFile(accountBalanceFile)

	if err != nil {
		fmt.Println("-------")
		fmt.Println("ERROR")
		fmt.Println(err)
		// Function for tracking the error infos
		// panic("Can't continue, sorry.")
		fmt.Println("-------")
	}

	for {
		presentOptions()

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
		fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
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
		fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
	}
	return accountBalance
}
