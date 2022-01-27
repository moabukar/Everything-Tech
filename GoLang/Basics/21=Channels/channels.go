// Implementation of Channels in Golang (Go Programming Language)

package main

import (
	"fmt"
	"sync"
)

// Create a struct that contains basic info
type Info struct {
	Name    string
	Address string
	About   string
}

// Declaring a Waitgroup
// A WaitGroup is used to wait for all the Goroutines launched to finish
var wg sync.WaitGroup

func main() {
	//Declaring a Channel
	channel := make(chan Info)
	wg.Add(2)

	// 
	go func() {
		// Sending data into the channel
		channel <- Info{
			Name:    "John",
			Address: "Ohio",
			About:   "Veteran :)",
		}

		channel <- Info{
			Name:    "Milinches",
			Address: "Arizona",
			About:   "GitHub MetaVerse.",
		}

		channel <- Info{
			Name: "Golang",
			Address: "NY City",
			About: "ListenandServe",
		}
		
		defer func ()  {
			close(channel)	
		}()
		
		// Tells the wait group, this particular goroutine has completed. (waitgroup - 1) 
		wg.Done()
	}()

	go func() {
		// receiving data from the channel
		for {
			// We probably have seen and used the switch statements in other languages
			// Bet you are shocked seeing a ``select`` statement been used with ``case``
			// ``Select`` statements in Golang are just like ``switch`` statement.
			// Select statement lets multiple Goroutine wait on multiple communication process.
			select {
			case rChan := <-channel:
				fmt.Printf("%s - lives at: %s - Boring profile: %s \n", rChan.Name, rChan.Address, rChan.About)
			}
			wg.Done()
		}
	}()

	wg.Wait()
}