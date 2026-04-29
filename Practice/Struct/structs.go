package main

import (
	"errors"
	"fmt"
	"time"
)

type user struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

// using constructor
func newUser(ufirstName, ulastName, ubirthdate string) (*user, error) {
	if ufirstName == "" || ulastName == "" || ubirthdate == "" {
		return nil, errors.New("firstName,LastName and dob is compulasory")
	}

	return &user{
		firstName: ufirstName,
		lastName:  ulastName,
		birthDate: ubirthdate,
		createdAt: time.Now(),
	}, nil
}

// method attaching stuct to user
func (u *user) outputUserDetails() {

	fmt.Println(u.firstName, u.lastName, u.birthDate)

}

func (u *user) clearUserDetail() {
	u.firstName = ""
	u.lastName = ""
}

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthDate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var userData *user

	userData, err := newUser(userFirstName, userLastName, userBirthDate)
	if err != nil {
		fmt.Print("erroe")
		return
	}

	userData.outputUserDetails()
	userData.clearUserDetail()
	userData.outputUserDetails()

}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
