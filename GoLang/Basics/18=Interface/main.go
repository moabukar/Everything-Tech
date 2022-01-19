package main

import "fmt"

type human struct {
	name string
}

type animal struct {
	name string
}

type furniture struct {
	name string
}

type testInterface interface {
	// types contain Speak() method
	Speak() string
}

func main() {
	// Interfaces are named collections of method signatures.
	//For our example we’ll implement this interface on human and animal types.
	irene := human{"irene"}
	cat := animal{"katty"}
	chair := furniture{"armless chair"}
	// human and animal type contain Speak()
	// method so can access to log() function

	log(irene)
	log(cat)

	// log(chair)
	// Error: cannot use chair (variable of type furniture)
	// as testInterface value in argument to log:
	// missing method Speak()
	fmt.Println(chair)
}

// func log(h human){
// 	fmt.Println(h.Speak())
// }
// func log(a animal) {
// 	fmt.Println(a.Speak())
// }
// the interfaces solving this big problem on the top
func log(in testInterface) {
	fmt.Println(in.Speak())
}

func (human) Speak() string {
	return "hi hi"
}

func (animal) Speak() string {
	return "meow meow"
}