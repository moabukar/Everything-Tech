# IMPLEMENTATION OF CHANNELS IN GOLANG

Channels provides the easuest way to allow `Goroutines` communicate with each other. Communications using channels are bidirectional by default - you can `send` and `recieve` data from the same channel.

Channels works with two operations which is `SENDING DATA` and `RECEIVING DATA`. 

## SENDING DATA:

The send operator is used to send data from one goroutine to another goroutine with the help of a channel.
```go 
    sendChan <- "Hello there"
```

## RECEIVING DATA

The receive operator is used to receive data that is being sent.
```go
    receiveChan := <-sendChan
```

## Example of Send and Receive Channels.

```go
    package channels

    import (
        "fmt"
        "sync"
    )

    var wg sync.WaitGroup

    func main() {
        // Creating the Channel
        sendChan := make(chan string)
        wg.Add(2)

        go func ()  {
            // Receiving data from the channel
            receiveChan := <-sendChan
            fmt.Println(receiveChan)
            wg.Done()
        }()

        go func ()  {
            // Sending data into the Channel
            sendChan <- "Implementation of a channel"
            wg.Done()
        }()

        wg.Wait()
    }
```