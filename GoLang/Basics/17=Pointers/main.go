//Go has pointers. A pointer holds the memory address of a value.
//A pointer is a variable that stores the address of a value,
// rather than the value itself. If you think of a computer’s memory (RAM)
// as a JSON object, a pointer would be like the key, and a normal variable
// would be the value.

package main

import "fmt"

func main() {
   var human string = "irene"   // actual variable declaration
   var female *string      // pointer variable declaration

   female = &human
   //return irene
   fmt.Println(human)
   
   // store address of a in pointer variable
   fmt.Printf("Address of a variable: %x\n", &human  )

   // address stored in pointer variable 
   fmt.Printf("Address stored in female variable: %x\n", *female )

   // access the value using the pointer
   fmt.Printf("Value of *female variable: %x\n", *female )

}
// &Variable means give me the memory address
	// *Pointer means give me the value this memory
	// address is pointing at

