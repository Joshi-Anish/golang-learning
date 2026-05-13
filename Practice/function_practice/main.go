package main

import (
	"fmt"
)

func main() {

	numbers := []int{2, 3, 4}
	doublev := transformNumber(&numbers, double)
	fmt.Println(doublev)
	triplev := transformNumber(&numbers, triple)
	fmt.Println(triplev)

}
func transformNumber(numbers *[]int, tranform func(int) int) []int {

	dNumber := []int{}

	for _, val := range *numbers {
		dNumber = append(dNumber, tranform(val))
	}
	return dNumber
}
func double(numbers int) int {
	return numbers * 2
}
func triple(number int) int {
	return number * 3
}
