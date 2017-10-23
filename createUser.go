package createUser

import "fmt"

type User struct {
	Name string
}

//  funtion available from createUser //
func PrintCreateUserName(userData User) {
	fmt.Println(userData.Name + "\n")
}

// method for the "User type" //
func (userData *User) GetName() string {
	return userData.Name
}
