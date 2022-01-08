package main

import "fmt"

// Go’s structs are typed collections of fields. 
//They’re useful for grouping data together to form records.
//structs are mutable
// This user struct type has id,first name, lastname,age
//email and password fields


type User struct {
	ID                string           
	FirstName         string                 
	LastName          string 
	Age               int                
	Email             string                
	Password          string 
}

func main() {
	// New user
	var newuser User

	// Set values
	newuser.ID = "12345678"
	newuser.FirstName = "irene"
	newuser.LastName = "ufia"
	newuser.Age = 25
	newuser.Email = "irenecufia@gmail.com"
	newuser.Password = "1234"

	fmt.Println(newuser.FirstName)

// or we can declare and set values with different
	// way...
	anotheruser := User{
		"56789",
		"jubril",
		"aminu",
		25,
		"jubrilaminu01@gmail.com",
		"1245",
	}

	fmt.Println(anotheruser)
}