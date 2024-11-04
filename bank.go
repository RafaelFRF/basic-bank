package main

import "fmt"

var accountBalance float64 = 1000.0

func main() {

	for {
		initialQuestions()

		var choice int = getUsersChoice()

		if choice == 1 {
			fmt.Println("Your balance is: ", accountBalance)
		} else if choice == 2 {
			depositAmount()
		} else if choice == 3 {
			withdrawAmount()
		} else {
			fmt.Println("Goodbye!")
			break
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

func withdrawAmount() {
	fmt.Print("Withdraw amount: ")
	var withdrawAmount float64
	fmt.Scan(&withdrawAmount)

	if withdrawAmount <= 0 {
		fmt.Println("Invalid amount. Must be greater than 0.")
		return
	}
	if withdrawAmount > accountBalance {
		fmt.Println("Invalid amount. You can't withdraw more than you have.")
		return
	}

	accountBalance -= withdrawAmount
	fmt.Println("Balance updated! New amount:", accountBalance)
}

func depositAmount() {
	fmt.Print("Your deposit: ")
	var depositAmount float64
	fmt.Scan(&depositAmount)

	if depositAmount <= 0 {
		fmt.Println("Invalid amount. Must be greater than 0.")
		return
	}

	accountBalance += depositAmount
	fmt.Println("Balance updated! New amount:", accountBalance)
}
