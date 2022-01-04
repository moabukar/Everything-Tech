package main

import "fmt"

func main() {
	// In Go, an array is a numbered sequence of
	// elements of a specific length.

	var a [4]int
	fmt.Println(a)
	// Here we create an array a that will hold
	// exactly 4 ints. The type of elements and
	// length are both part of the array’s type.
	// By default an array is zero-valued, which
	// for ints means 0s.

	a[0] = 154
	fmt.Println(a)
	// We can set a value at an index using the
	// array[index] = value syntax,
	// and get a value with array[index].
	fmt.Println(a[0])

	// The builtin len returns the length of an array
	fmt.Println(len(a))

	// second way for declaring arrays
	b := [4]int{}
	fmt.Println(b)

	// or
	c := [4]int{10, 20, 30, 40}
	fmt.Println(c)

	// or
	d := [2][2]string{}
	d[0][0] = "irene"
	d[0][1] = "ufia"
	// d[0][2] = "ufia" return the error becuase index
	// 2 in not declare
	fmt.Println(d)
}