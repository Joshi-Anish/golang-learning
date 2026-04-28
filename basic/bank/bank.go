package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

// universal cosntant
const accountBalanceFile = "balance.txt"

// writing in a file xalled balance.txt
func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)

}

// reading a file
func readBalanceFromFile() (float64, error) {

	data, err := os.ReadFile(accountBalanceFile)

	if err != nil {
		return 1000, errors.New("Failed to find balance file")
	}

	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)

	if err != nil {
		return 1000, errors.New("conversion failed")
	}
	return balance, nil
}

func main() {
	//variabble declaration
	Balance, err := readBalanceFromFile()

	if err != nil {
		fmt.Println("Error")
		fmt.Print(err)
		fmt.Println("---------")
	}
	fmt.Println("------Welcome to GO bank-----------")

	//infinite loop break lae loop end garxa
	for {
		fmt.Println("What do you want to do?")
		fmt.Println("1.Check Balance")
		fmt.Println("2.Deposit money")
		fmt.Println("3.Withdeaw money")
		fmt.Println("4.Exit")

		var choice int
		fmt.Scan(&choice)
		fmt.Println("your choice is:", choice)

		if choice == 1 {
			fmt.Println("Your balance is :", Balance)
			writeBalanceToFile(Balance)
		} else if choice == 2 {

			var deposit float64
			fmt.Print("How much you want to deposit")
			fmt.Scan(&deposit)

			if deposit <= 0 {
				fmt.Print("Must deposit more than 0")
				continue
			}
			fmt.Println("you deposited:", deposit)
			Balance += deposit
			fmt.Println("you New balance is:", Balance)
			writeBalanceToFile(Balance)

		} else if choice == 3 {
			var withdraw float64
			fmt.Println("How much you want to withdraw :")
			fmt.Scan(&withdraw)

			if withdraw <= 0 {
				fmt.Print("Enter amount more than 0")
				continue
			}
			if withdraw > Balance {
				fmt.Println("Insufficent amount")
				continue
			}

			Balance -= withdraw
			fmt.Println("your new Balance is :", Balance)
			writeBalanceToFile(Balance)

		} else {
			println("goodbye")
			break
		}

	}
	fmt.Print("Thnaks for choising our bank")
}
