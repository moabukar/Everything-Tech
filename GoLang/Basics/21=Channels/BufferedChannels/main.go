// Buffered Channels
// By default, channels are unbuffered. Which implies that the channel
// would only accept data into the channel if there is a corresponding receive.
// Buffered Channels are just like internal data stores, they allow a limited
// number of values without a reciver for thos values. Buffered Channels are blocked
// when the channel is full.

// We'll see how to create a buffered Channel below.

package main

import "fmt"

type Info struct {
	Name    string
	Address string
}

func main() {
	// We just made a buffered channel by adding an extra parameter to the make function
	// The extra parameter we added specifies the size of the buffer
	// We created a buffered that can
	bufChannel := make(chan Info, 2)
	bufChannel <- Info{
		Name: "John",
		Address: "Georgia",
	}
	bufChannel <- Info{
		Name: "James",
		Address: "NY City",
	}

	// Since the size of the buffered channel is ``2``, sending another data to the
	// channels would result in a deadlock. What is a deadlock?
	// Deadlock is a situation where a set of processes are blocked because each process
	// is holiding a resource and waiting for another resource acquired by another resource.

	// prints out the output
	fmt.Println(<-bufChannel)
	fmt.Println(<-bufChannel)
}
