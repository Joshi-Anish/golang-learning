package main

import (
	"fmt"
)

func main() {

	//like here we used [o]value and doesnt used 1 2 3 value of array
	productName := [4]string{"anish"}
	prices := [3]float64{2.2, 3.0, 12.2}
	fmt.Println(prices)

	//we used the 3rd vakue of array
	productName[2] = "joshi"
	fmt.Println(productName)
}
