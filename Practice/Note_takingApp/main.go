package main

import (
	"errors"
	"fmt"
)

func main() {

}

func getNoteData() (string, string, error) {
	title, err := getUserInput("Enter title:")

	if err != nil {
		fmt.Println("Error:input the title")
		return "", "", err
	}
	content, err := getUserInput("Enter the content:")

	if err != nil {
		fmt.Println("Error:input the content")
		return "", "", err
	}
	return title, content, err
}
func getUserInput(prompt string) (string, error) {
	fmt.Print(prompt)
	var value string
	fmt.Scan(&value)

	if value == "" {
		return "", errors.New("enter the text")
	}

	return value, nil
}
