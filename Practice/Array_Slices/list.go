package main

import (
	"fmt"
)

// dynamic array
func main() {
	prices := []float64{2.2, 1, 0}
	fmt.Println(prices)

	//updating this slice dynamically

	prices = append(prices, 3.3)
	fmt.Println(prices)
}

// func main() {

// 	//like here we used [o]value and doesnt used 1 2 3 value of array
// 	productName := [4]string{"anish"}
// 	prices := [3]float64{2.2, 3.0, 12.2}
// 	fmt.Println(prices)

// 	//we used the 3rd vakue of array
// 	productName[2] = "joshi"
// 	fmt.Println(productName)

// 	//slices
// 	newPrices := prices[1:3]
// 	fmt.Println(newPrices)

// 	//slices inside slices
// 	hightedPrice := newPrices[:1]
// 	fmt.Println(hightedPrice)
// }
