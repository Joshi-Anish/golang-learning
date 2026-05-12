package list

import (
	"fmt"
)

func main() {
	prices := []float64{2.2, 1, 0}
	fmt.Println(prices[:1])
	prices[1] = 9.9

	//updating this slice dynamically

	prices = append(prices, 5.99, 3.3, 12.99, 100.10)
	prices = prices[1:]
	fmt.Println(prices)

	discountedPrices := []float64{101.99, 80.99, 20.59}
	prices = append(prices, discountedPrices...)
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
