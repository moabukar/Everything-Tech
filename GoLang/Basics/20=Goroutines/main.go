package main

import (
	"fmt"
	"net/http"
)

func main() {
	// A goroutine is a lightweight thread managed
	// by the Go runtime.

	// Example 1
	links := []string{
		"http://google.com",
		"http://golang.org",
		"http://github.com",
		"http://stackoverflow.com",
	}
	//
	//for _ , link := range links {
	//	Example(link)
	//}
	// program is currently fetching one url at
	// the time waiting for that request complete
	// and then move on the next url
	// the impact of this is that if we have many
	// url and we want to fetch in our program
	// it'll take a long time to fetch every single one
	// because we're doing each one at the time
	// we want to fetch multiple request at the same
	// time actually how to make these request at the same
	// time?   GoRoutines is the answer


	// when we launch the program golang automatically
	// create the MAIN ROUTINE
	// when the main routine reach to the first
	// line of ExampleOne func that wait for complete
	// then get to the next line of code
	// this bad because take the long time to fetching
	// one url at the time
	// we can solve this problem with create a new
	// child routine by `go` keyword
	// how to do?
	// if we want to create a new child routine
	// we should use `go` keyword before calling function
	// example:   go Example(link)

	//for _ , link := range links {
	//	go Example(link)
	//}

	// again we have the new problem
	// because the main routine end up faster than
	// other routine, and we can't see the result of these
	// child routines
	// golang has the solution for this big problem
	// CHANNELS

	// how they work?
	// look at the GoRoutines.pptx slide 4

	// we'll create a channel that we can use to say
	// hey has this routine finished, has this finished
	// and has this finished ok then the main routine
	// can exit.

	// Create channel syntax
	channel := make(chan string)
	// channel string


	for _ , link := range links {
		go Example(link , channel)
	}

	// listening to the channel
	for i := 0 ; i < len(links) ; i++ {
		fmt.Println(<- channel)
	}
}

// Example
// receiving the channel as an argument
func Example(link string , channel chan string) {
	_ , err := http.Get(link)

	if err != nil{
		// We will talk about nil in a few
		// chapters later
		fmt.Println(link , "is down")
		// sending message to channel with this `<-` syntax
		channel <- "is down"
		return
	}

	fmt.Println(link , "is up")
	channel <- "is up"
}