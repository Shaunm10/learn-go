package main

import (
	"fmt"
	"os"
	"strconv"
)

const DataFileName = "balance.data"

/*
	GOALS:
	=================
	- Ask the user what operation they would like to perform in our bank application.
	"
	Welcome to Go Bank!
	What do you want to do?
	1. Check Balance
	2. Deposit money
	3. Withdraw money
	4. Exit
	Your choice:
	"
*/

func main() {
	var accountBalance = readBalanceToFile()
	var runApplicationLoop = true

	// show welcome message
	fmt.Println("******************")
	fmt.Println("Welcome to Go Bank!")

	for runApplicationLoop {

		menuChoice := getMainMenuSelection()

		switch menuChoice {
		case 1:
			showAccountBalance(accountBalance)
		case 2:
			handleDeposit(&accountBalance)
		case 3:
			handleWithdraw(&accountBalance)
		case 4:
			handleExit(accountBalance)
			runApplicationLoop = false
		default:
			fmt.Printf("%v is an invalid selection. Please try again.", menuChoice)
		}
	}

	fmt.Println("Shutting down...")
}

func getMainMenuSelection() (menuChoice int) {
	fmt.Println("")
	fmt.Println("What do you want to do?")
	fmt.Println("1. Check Balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")
	fmt.Println("Your choice:")
	fmt.Scan(&menuChoice)
	return
}

func showAccountBalance(accountBalance float64) {
	fmt.Printf("Your account balance is: $%.2f\n", accountBalance)
}

func handleExit(accountBalance float64) {
	writeBalanceToFile(accountBalance)
	fmt.Println("Goodbye!")
}

func handleDeposit(accountBalance *float64) {
	var depositAmount float64
	fmt.Println("How much would you like to deposit?")
	fmt.Scan(&depositAmount)

	if depositAmount < 0 {
		fmt.Println("Negative amounts are not permitted.")
		handleDeposit(accountBalance)
	} else {

		*accountBalance += depositAmount

		fmt.Printf("Your account balance is updated to $%.2f\n ", *accountBalance)
		return
	}
}

func handleWithdraw(accountBalance *float64) {
	var withdrawAmount float64
	fmt.Println("How much would you like to withdraw?")
	fmt.Scan(&withdrawAmount)
	if withdrawAmount > *accountBalance {
		fmt.Printf("You will have overdrawn you account, your current balance is $%.2f\n", *accountBalance)
		handleWithdraw(accountBalance)
	} else {
		*accountBalance -= withdrawAmount
		fmt.Printf("Your account balance is updated to $%.2f\n ", *accountBalance)
		return
	}
}

func writeBalanceToFile(accountBalance float64) {
	balanceText := fmt.Sprint(accountBalance)
	os.WriteFile(DataFileName, []byte(balanceText), 0644)
}

func readBalanceToFile() float64 {

	// TODO: handle error
	data, _ := os.ReadFile(DataFileName)
	balanceText := string(data)

	// TODO: handle error
	balance, _ := strconv.ParseFloat(balanceText, 64)
	return balance

}
