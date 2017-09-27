package createUser

import "fmt"

type User struct {
	Name string
}

func PrintCreateUserName(userData User) {
	fmt.Print(userData.Name)
}

func (userData *User) GetName() string {
	return userData.Name
}
