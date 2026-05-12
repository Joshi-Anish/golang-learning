package main

import (
	"fmt"
)

func main() {
	//1
	hobbies := [3]string{"playingCricket", "watchingShows", "bikeRiding"}
	fmt.Println("All hobbies", hobbies)

	//2
	fmt.Println("The first hobby", hobbies[0])

	newHobbies := hobbies[1:]
	fmt.Println("the 2 other hobbies", newHobbies)

	//3
	startingHobbies := hobbies[:2]
	differentWayStatingHobbies := hobbies[0:2]

	fmt.Println(startingHobbies)
	fmt.Println(differentWayStatingHobbies)

	//4
	secondHobby := startingHobbies[1:]
	fmt.Println(secondHobby)

	//5
	courseGoal := []string{"1.learn basic of go.", "2.can start doing adavance task"}
	fmt.Println(courseGoal)
	courseGoal[1] = "Do project"
	fmt.Println(courseGoal)
	courseGoal = append(courseGoal, "get a certificate")
	fmt.Println(courseGoal)
}

// Time to practice what you learned!

// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.
// 2) Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list
// 3) Create a slice based on the first element that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end)
// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.
