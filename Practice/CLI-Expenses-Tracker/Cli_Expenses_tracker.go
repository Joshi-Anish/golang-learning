package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func saveExpensesToFile(expenses float64) {
	data := fmt.Sprint(expenses)
	err := os.WriteFile("totalExpense.txt", []byte(data), 0644)

	if err != nil {
		fmt.Println("Error saving file")
		return
	}
	fmt.Println("Expenses is updated in FIle")
}

func readExpensesFromFile() (float64, error) {

	data, err := os.ReadFile("totalExpense.txt")

	if err != nil {
		return 0.0, errors.New("No file exits")
	}
	expensesText := string(data)

	expenses, err := strconv.ParseFloat(expensesText, 64)

	if err != nil {
		return 0.0, errors.New("Error while reading")
	}

	return expenses, nil
}

func main() {
	fmt.Print("Welcome to CLI Expense Tracker")

	expenses, err := readExpensesFromFile()

	// If file does not exist
	if err != nil {
		fmt.Println(err)
		expenses = 0.0
	}
	//variable declartion
	// var addExpenses float64
	// var viewExpenses float64
	// var totalExpenses float64

	for {
		fmt.Println("what do you want to do?")
		fmt.Println("1.Add Expenses")
		fmt.Println("2.view Expenses")
		fmt.Println("3.Show total Expenses")
		fmt.Println("4.exit")

		var choice int
		fmt.Scan(&choice)

		switch choice {

		case 1:
			fmt.Println("Add Expense")
			fmt.Println("Enter expenses name:")
			var expensesName string
			fmt.Scan(&expensesName)

			fmt.Println("Enter the expense amount")
			var expenseAmount float64
			fmt.Scan(&expenseAmount)

			expenses += expenseAmount

			fmt.Println(expensesName, "added successfully!")

			// Save total expenses to file
			saveExpensesToFile(expenses)

		case 2:
			fmt.Println("This will be added later")

		case 3:
			fmt.Println(" Total Expenses:", expenses)

		case 4:
			fmt.Println("Exiting.....")
			return

		default:
			fmt.Println("Invalid choice")
			return
		}

	}

}
