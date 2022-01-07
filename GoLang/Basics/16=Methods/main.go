package main

import "fmt"

// Go does not have classes. However, go supports 
// methods on struct types
type user []string


// A method is a function with a special receiver argument.
//below are methods with a receiver of user type
func (u user) logger() {
	fmt.Println(u)
}

func (u user) customLogger(from, to int) {
	fmt.Println(u[from:to])
}

func (u user) addUser(name string) user {
	list := u
	list = append(list, name)

	return list
}

//main function
func main() {
	newUser := user{"irene", "jubril", "abubakar", "dod"}

	// fmt.Println(newUser)
	newUser.logger()
	newUser.customLogger(1, 4)

	addedUser := newUser.addUser("king")
	addedUser.logger()

}
