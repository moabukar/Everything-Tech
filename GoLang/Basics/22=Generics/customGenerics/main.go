// Writing custom generics in Golang
package main

import "fmt"

// Create a program that adds up all the numbers in an Array

// Creating a custome Generics.
// Our customGenerics only supports these data types.
type customGenerics interface {
	int32 | int64 | uint32 | uint64 | float64
}

func addNumbers[T customGenerics](arrayVal []T) T {
	var output T = 0

	for _, value := range arrayVal {
		output += value
	}
	return T(output)
}

func main() {
	// output
	fmt.Println(addNumbers([]int32{-1, 1, 2, 3, 4, 5, 6}))
	fmt.Println(addNumbers([]int64{-1, 1, 2, 3, 4, 5, 6}))
	fmt.Println(addNumbers([]uint32{1, 2, 3, 4, 5, 6}))
	fmt.Println(addNumbers([]uint64{1, 2, 3, 4, 5, 6}))
	fmt.Println(addNumbers([]float64{1.11, 2.1, 3.4, 5.6, 5.4, 6.6}))
}