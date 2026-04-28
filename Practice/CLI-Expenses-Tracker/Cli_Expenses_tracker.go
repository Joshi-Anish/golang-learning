package main

import (
	"fmt"
)

func main() {
	fmt.Print("Welcome to CLI Expense Tracker")

	var expenses = 0.0
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

		case 2:
			fmt.Println("This will be added later")

		case 3:
			fmt.Println(" Total Expenses:", expenses)

		case 4:
			fmt.Println("Exiting...")
			return

		default:
			fmt.Println("Invalid choice")
			return
		}

	}

}
