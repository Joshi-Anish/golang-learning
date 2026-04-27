package main

import (
	"fmt"
)

func main() {
	fmt.Print("---------Welcome to Profit Calculator----\n")

	//variable declaration
	var revenue float64
	var expenses float64
	var taxrate float64

	//user input
	fmt.Print("Enter the revemue amount:")
	fmt.Scan(&revenue)

	fmt.Print("Enter the expenses amount:")
	fmt.Scan(&expenses)

	fmt.Print("Enter the tax rate:")
	fmt.Scan(&taxrate)

	//calculation
	//EBT=Earning before tax , profit=Earining after tax

	var EBT = revenue - expenses
	fmt.Println("Value of before tax is", EBT)

	var profit = (revenue - expenses) * (1 - (taxrate / 100))
	fmt.Println("Value after tax is ", profit)

	var ratio = EBT / profit
	fmt.Println("the ratio of value before and ater tax is", ratio)
}
