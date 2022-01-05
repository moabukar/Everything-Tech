package main

import "fmt"

// Data types in golang
func main() {
	// Integer types
	// Default value for integer types = 0
	var age int = 25
	var x int
	fmt.Println(x)
	// x variable return the 0

	// Floating types
	// Default value for float types = 0
	var pi float64 = 3.14

	// String types
	// Default value for string types = ""
	var name string = "irene"
	var family string
	fmt.Println(family)
	// family variable return the ""

	// Boolean types
	// Default value for boolean types = false
	var human bool = true
	var toggle bool
	fmt.Println(toggle)
	// toggle variable return the false value

	fmt.Println("my age: ", age)
	fmt.Println("my name: ", name)
	fmt.Println("im a human? ", human)
	fmt.Println("pi number: ", pi)
}