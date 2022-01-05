package main

import "fmt"

func main() {
	//there are two ways to declare a variable in go language

	// first
	var name string = "irene"
	var age = 25

	// second
	nationality := "nigerian"

	fmt.Println("my name is: ", name)
	fmt.Println("my age: ", age)
	fmt.Println("and my nationality: ", nationality)
}

//You can also initialize variable on the same line

/*var (
name, nationality, age = "Irene ufia", "nigerian", 25
)*/
