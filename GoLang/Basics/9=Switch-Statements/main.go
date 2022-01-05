package main

import "fmt"

func main() {
	// Switch statements express conditionals across
	// many branches.

	x := 2

	switch x {
	case 0:
		fmt.Println("x variable === 0")

	case 1:
		fmt.Println("x variable === 1")

	case 2:
		fmt.Println("x variable === 2")
	}

}