package main

import (
	"fmt"
)

func main() {

	numbers := []int{2, 3, 4}
	double := doubleNumber(&numbers)
	fmt.Println(double)

}
func doubleNumber(numbers *[]int) []int {

	dNumber := []int{}

	for _, value := range *numbers {
		dNumber = append(dNumber, value*2)
	}
	return dNumber
}
