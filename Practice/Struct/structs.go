package main

import (
	"fmt"
	"struct.com/struct/user"
)

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthDate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var userData *user.User

	userData, err := user.New(userFirstName, userLastName, userBirthDate)

	if err != nil {
		fmt.Println("error")
		return
	}

	userData.OutputUserDetails()
	userData.ClearUserDetail()
	userData.OutputUserDetails()

}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
