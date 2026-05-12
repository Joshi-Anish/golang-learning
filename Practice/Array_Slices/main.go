package main

import (
	"fmt"
)

func main() {
	username := make([]string, 2, 5)
	//2 means front ma 2 empty slot hunxa and 5 vanae ko memeory capactiy allocate gardiya ko
	username[0] = "anish"
	username = append(username, "joshi")

	fmt.Println(username)

	for index, value := range username {
		fmt.Println(index)
		fmt.Println(value)
	}
}
