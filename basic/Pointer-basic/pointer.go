package main

import (
	"fmt"
)

func main() {
	age := 32 //regular variable

	userAge := &age //pointer which store age address

	fmt.Println("Age address", userAge)

	adultYear := getAdultYears(userAge)
	fmt.Println(adultYear)

}

// we have called pointer and used it . so pointer cha vanae use * to get value of that address
func getAdultYears(userAge *int) int {
	return *userAge - 18
}
