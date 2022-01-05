package main

import "fmt"

func main() {
	/* Slices are a key data type in Go, giving a more powerful interface to sequences
	than arrays.Thoug it is just like an array having an index value and length,
	the difference between a slice and an array is that, sixe of slices can be resized but arrays are fixed-sized*/

	//declaring
	slice := []int{}
	fmt.Println(slice)

	// Unlike arrays, slices are typed only by the
	// elements they contain (not the number of elements)

	// append a value
	slice = append(slice, 10)
	fmt.Println(slice)
	// slice[1] = 20 return the error because we not
	// define a start length

	sliceTwo := make([]string, 2)
	sliceTwo[0] = "irene"
	sliceTwo[1] = "ufia"
	// sliceTwo[2] = "nothing" return the error because
	// we defined a 2 length and we can just appended to
	// the slice
	sliceTwo = append(sliceTwo, "dom")
	fmt.Println(sliceTwo)

	// basic example
	x := []int{}

	for i := 0; i < 5; i++ {
		x = append(x, i)
	}

	fmt.Println(x)
}
