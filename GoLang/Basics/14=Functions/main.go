package main

import "fmt"

func main() {
	fmt.Println(sum(1, 2))
	// return 3

	users := users("irene", "ufia")
	fmt.Println(users)
	// return ["irene" , "ufia"]

	// Closure functions
	// GO functions can contain functions
	hello := func(name string) {
		fmt.Println("hello", name)
	}
	hello("irene")
	// return hello irene

	name, age := multipleReturns("irene", 25)
	fmt.Println(name, age)
	// return irene , 25

	fake := falseFunc(1)
	fmt.Println(fake)
	// return 2
}

// Here’s a function that takes two ints and
// returns their sum as an int.
func sum(a, b int) int {
	return a + b
}

// or
// func sum(a int , b int) int {}

// Variadic functions can be called with any number of
// trailing arguments for example:
func users(names ...string) []string {
	return names
}

// Multiple Return Values
func multipleReturns(name string, age int) (string, int) {
	return name, age
}

// Recursive functions
func falseFunc(x int) int {
	if x == 0 {
		return 1
	}

	return x + falseFunc(x-1)
}