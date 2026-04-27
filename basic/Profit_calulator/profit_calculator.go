// package main

// import (
// 	"fmt"
// )

// func main() {
// 	fmt.Print("---------Welcome to Profit Calculator----\n")

// 	//variable declaration
// 	var revenue float64
// 	var expenses float64
// 	var taxrate float64

// 	//user input
// 	fmt.Print("Enter the revemue amount:")
// 	fmt.Scan(&revenue)

// 	fmt.Print("Enter the expenses amount:")
// 	fmt.Scan(&expenses)

// 	fmt.Print("Enter the tax rate:")
// 	fmt.Scan(&taxrate)

// 	//calculation
// 	//EBT=Earning before tax , profit=Earining after tax

// 	var EBT = revenue - expenses
// 	fmt.Println("Value of before tax is", EBT)

// 	var profit = (revenue - expenses) * (1 - (taxrate / 100))
// 	fmt.Println("Value after tax is ", profit)

// 	var ratio = EBT / profit
// 	fmt.Println("the ratio of value before and ater tax is", ratio)
// }

// using function
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
	// there is & here because we are sending value to function
	inputData(&revenue, &expenses, &taxrate)

	//output and calulation function call

	//like there is no & here because we are recving value from fuction using return
	EBT, profit, ratio := calc(revenue, expenses, taxrate)
	fmt.Println("Value of before tax is", EBT)
	fmt.Println("Value after tax is ", profit)
	fmt.Println("the ratio of value before and ater tax is", ratio)

}

func inputData(revenue, expenses, taxrate *float64) {

	fmt.Print("Enter the revemue amount:")
	fmt.Scan(revenue)

	fmt.Print("Enter the expenses amount:")
	fmt.Scan(expenses)

	fmt.Print("Enter the tax rate:")
	fmt.Scan(taxrate)

}

func calc(revenue, expenses, taxrate float64) (float64, float64, float64) {
	var EBT = revenue - expenses

	var profit = (revenue - expenses) * (1 - (taxrate / 100))

	var ratio = EBT / profit

	return EBT, profit, ratio

}
