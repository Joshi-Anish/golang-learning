package main

import (
	"fmt"
	"math"
)

func main() {

	var intialInvestment float64
	var expectedReturnRate float64
	var years int

	const inflation = 6.5

	fmt.Print("Welcome to Investment Calclator\n")
	//user input
	fmt.Print("enter intial Investment:")
	fmt.Scan(&intialInvestment)

	fmt.Print("enter expected return rate:")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("enter total no years:")
	fmt.Scan(&years)

	var futureValue = float64(intialInvestment) * math.Pow(1+expectedReturnRate/100, float64(years))
	fmt.Println(futureValue)

	var inflationValue = futureValue / math.Pow(1+inflation/100, float64(years))
	fmt.Println(inflationValue)
}
