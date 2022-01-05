package main

import "fmt"

func main() {
	// Maps are Go’s built-in associative data
	// type (sometimes called hashes or dicts in
	// other languages).

	// To create an empty map, use the builtin make:
	// make(map[key-type]val-type).
	m := make(map[int]string)
	m[1] = "irene"
	m[2] = "ufia"
	fmt.Println(m)

	// The builtin delete removes key/value pairs from a map.
	delete(m, 1)
	fmt.Println(m)

	// The optional second return value when getting
	// a value from a map indicates if the key was
	// present in the map. This can be used to
	// disambiguate between missing keys and keys
	// with zero values like 0 or "".
	_, value := m[4]
	if !value {
		fmt.Println("not found!")
	} else {
		fmt.Println("true")
	}

	/* You can also declare and initialize a new map in
	the same line with this syntax.*/
	second := map[int]string{1: "irene"}
	fmt.Println(second)
}
