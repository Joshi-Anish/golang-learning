package main

import (
	"fmt"
	"math"
)

func main() {

	var intialInvestment = 1000
	var expectedReturnRate = 5.5
	var years = 10

	const inflation = 6.5

	var futureValue = float64(intialInvestment) * math.Pow(1+expectedReturnRate/100, float64(years))

	fmt.Println(futureValue)

	var inflationValue = futureValue / math.Pow(1+inflation/100, float64(years))

	fmt.Println(inflationValue)
}
