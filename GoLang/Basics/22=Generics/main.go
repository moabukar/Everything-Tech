// Generics in Golang
//
// Implementing a Binary search tree using Golang
package main

import (
	"constraints"
	"fmt"
	"math"
	"sort"
)

// This new BinaryTree function looks different from the functions we've been writing.
//
// What Changed?
//
/*
	func BinaryTree[T constrainst.Ordered](searchArr []T, searchValue T) 
*/
// Well, that is Generics in Golang. 
//
//``constraints.Ordered`` permits any ordered type
// that supports the operators < <= => >
//
// Below are currently what the ``constraints.Ordered`` supports.
// You can check out the constraints Documentation for more
/* 
	type Ordered interface {
		Integer | Float | ~string
	}
*/
func BinaryTree[T constraints.Ordered](searchArr []T, searchValue T) bool {
	// Array needs to be sorted. Since we are using generics, we will have to use use the ``sort.Slice`` function
	sort.Slice(searchArr, func(i, j int) bool {
		return searchArr[i] < searchArr[j]
	})

	min := 0
	max := len(searchArr) - 1 // The length of the array - 1

	
	// This ``for`` loop in most languages are what we call the ``while`` loop
	for {
		splitArr := math.Floor(float64(min+max) / 2)
		convSplit := int(splitArr)

		if max < min {
			return false
		}

		if searchArr[convSplit] == searchValue {
			return true
		}

		if searchArr[convSplit] < searchValue {
			min = convSplit + 1
		} else {
			max = convSplit - 1
		}
	}
}

func main() {
	// output
	fmt.Println(BinaryTree([]string{"John", "James", "Jack", "Jake", "Jamal"}, "James")) // returns true
	fmt.Println(BinaryTree([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 9))                         // returns true
	fmt.Println(BinaryTree([]float64{1.0, 2.0, 3.0, 4.6, 5.7, 6.3, 7.7, 8.8, 9.1}, 9.0)) // returns false
}
