package main

import "fmt"

func main() {
	// range iterates over elements in a variety
	// of data structures

	nums := [4]int{1, 2, 3, 4}
	sum := 0
	for _, i := range nums {
		sum += i
	}
	fmt.Println(sum)
	// range on arrays and slices provides both
	// the index and value for each entry.
	// Above we didn’t need the index, so we ignored
	// it with the blank identifier _.

	// range on map iterates over key/value pairs.
	m := map[int]string{1: "irene", 2: "ufia"}
	for key, value := range m {
		fmt.Println(key, value)
	}

	// range on strings iterates over Unicode code points.
	str := "irene"
	for index, uni := range str {
		fmt.Println(index, uni)
	}
}