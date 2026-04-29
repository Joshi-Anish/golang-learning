package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

// using constructor
func NewUser(ufirstName, ulastName, ubirthdate string) (*User, error) {
	if ufirstName == "" || ulastName == "" || ubirthdate == "" {
		return nil, errors.New("firstName,LastName and dob is compulasory")
	}

	return &User{
		firstName: ufirstName,
		lastName:  ulastName,
		birthDate: ubirthdate,
		createdAt: time.Now(),
	}, nil
}

// method attaching stuct to user
func (u *User) OutputUserDetails() {

	fmt.Println(u.firstName, u.lastName, u.birthDate)

}

func (u *User) ClearUserDetail() {
	u.firstName = ""
	u.lastName = ""
}
