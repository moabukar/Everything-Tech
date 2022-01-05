package main

import "fmt"

// Here are some basic types of for loops, note that for is 
//the only looping construct in go

func main() {
	
	// The most basic type, with a single condition.
	i := 0
	for i <= 5 {
		fmt.Println(i)
		i++
	}

	// The classic initial/condition/update for loop
	for j := 0; j <= 10; j++ {
		fmt.Println(j)
	}

	// for without a condition will loop repeatedly
	// until you break out of the loop or return from
	// the enclosing function.
	for {
		fmt.Println("for without a condition")
		break
		// or return for break out the loop
	}
}