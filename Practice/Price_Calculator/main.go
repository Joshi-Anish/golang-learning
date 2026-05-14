package main

import "fmt"

func main() {
	prices := []float64{10, 20, 30}
	taxRate := []float64{0, 0.07, 0.1, 0.15}

	result := make(map[float64][]float64)

	for _, taxRate := range taxRate {
		taxIncludedPrice := make([]float64, len(prices))

		for priceIndex, prices := range prices {
			taxIncludedPrice[priceIndex] = prices * (1 + taxRate)
		}
		result[taxRate] = taxIncludedPrice

	}
	fmt.Println(result)

}
