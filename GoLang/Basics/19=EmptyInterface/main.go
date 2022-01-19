package main

import "fmt"

func main () {
	// The interface type that specifies zero methods
	// is known as the empty interface:
	// interface{}

	// An empty interface may hold values of any type.
	// (Every type implements at least zero methods.)
	// Empty interfaces are used by code that handles
	// values of unknown type

	unknownType(20)
	unknownType("hi")
	unknownType([]string{"irene" , "mary"})
}

// Example
// unknown type
func unknownType(i interface{}) {
	switch theType := i.(type) {
	case int:
		fmt.Println("int")

	case string:
		fmt.Println("string")

	default:
		fmt.Println(theType)
	}
}