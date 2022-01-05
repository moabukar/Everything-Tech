package main

import "fmt"

const DATABASE_URL = "mongodb://localhost:25408"

func main() {
	// Go supports constants of character, string,
	// boolean, and numeric values.

	fmt.Println(DATABASE_URL)

	const pi = 3.14
	fmt.Println(pi)
}