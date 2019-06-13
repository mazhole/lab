package main

import (
    "fmt"
    "math/rand"
    "time"
)

const messageCount = 10


func main() {
    fmt.Println("Example 1")

    // Creates channel (reference type)
    // Created channel is unbuffered i.e. without specified of buffer size
    // Goroutine-provider can't write to unbuffered channel until goroutine-consumer starts read from it
    c1 := make(chan string)

    go provider(c1)

    for i := 0; i < messageCount; i++ {
        fmt.Println("=> chan consumer reads data:", <- c1)
    }

    fmt.Println("\nExample 2")

    // Creates channel (reference type)
    // Created channel is buffered i.e. with specified of buffer size
    // Goroutine-provider can write to buffered channel until channel size fills up
    // Additional messages is written, if the consumer read some messages from the channel 
    c2 := make(chan int, messageCount)

    go func(c chan <- int){ 
        for i := 0; i < messageCount + 2; i++ {
            c <- i
            time.Sleep(1 * time.Second)
        }
        close(c)
    }(c2)

    c2 <- 100

    // Expected 166
    fmt.Printf("=> Result: %d\n", consumer(c2))
    // Shows the result after 10 seconds
    //=> Result: 166

    // Deadlock!
    // Documentation: https://golang.org/doc/effective_go.html#channels
    // If the channel is unbuffered, the sender blocks until the receiver has
    // received the value. If the channel has a buffer, the sender blocks only
    // until the value has been copied to the buffer; if the buffer is full, this
    // means waiting until some receiver has retrieved a value.
    //c := make(chan int)    
    //c <- 1   
    //fmt.Println(<-c)

    fmt.Println("\nExample 3")

    c3 := channelProvider("Hello")
    for i := 0; i < messageCount; i++ {
        fmt.Println("=>", <-c3)
    }

    fmt.Println("\nExample 4")

    c4 := channelProvider("Hello")
    c5 := channelProvider("Bye")
    c6 := multiChannel(c4, c5)
    for i := 0; i < messageCount * 2; i++ {
        fmt.Println("=>", <-c6)
    }
}

// Writes data into the channel
// 'chan <-' means that the channel is write only
// If try to read data from the channel in bounds of this method,
// compiler error: invalid operation: <-c (receive from send-only type chan<- string)
func provider(c chan <- string) {
    for i := 0; i < messageCount; i++ {
        c <- fmt.Sprintf("chan provider info %d", i)
    }
}

// Reads data from the channel
// '<- chan' means that the channel is read only
// If try to write data into the channel in bounds of this method,
// compiler error: invalid operation: c <- 4 (send to receive-only type <-chan int)
func consumer(c <- chan int) (result int) {
    for i := range c {
        fmt.Println("=> chan consumer reads item from the channel:", i)
        result += i
    }
    return
}

func multiChannel(c1, c2 <- chan string) <-chan string {
    c := make(chan string)
    go func() {
        // If comment 'for', deadlock exists
        for {
            c <- <- c1
        }
    }()
    go func() {
        // If comment 'for', deadlock exists
        for {
            c <- <- c2
        }
    }()
    return c
}

func channelProvider(info string) <-chan string {
    c := make(chan string)
    go func() {
        for i := 0; i < messageCount; i++ {
            c <- fmt.Sprintf("Message %s %d", info, i)
            time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
        }
    }()
    return c
}

//Output:
//Example 1
//=> chan consumer reads data: chan provider info 0
//=> chan consumer reads data: chan provider info 1
//=> chan consumer reads data: chan provider info 2
//=> chan consumer reads data: chan provider info 3
//=> chan consumer reads data: chan provider info 4
//=> chan consumer reads data: chan provider info 5
//=> chan consumer reads data: chan provider info 6
//=> chan consumer reads data: chan provider info 7
//=> chan consumer reads data: chan provider info 8
//=> chan consumer reads data: chan provider info 9
//
//Example 2
//=> chan consumer reads item from the channel: 100
//=> chan consumer reads item from the channel: 0
//=> chan consumer reads item from the channel: 1
//=> chan consumer reads item from the channel: 2
//=> chan consumer reads item from the channel: 3
//=> chan consumer reads item from the channel: 4
//=> chan consumer reads item from the channel: 5
//=> chan consumer reads item from the channel: 6
//=> chan consumer reads item from the channel: 7
//=> chan consumer reads item from the channel: 8
//=> chan consumer reads item from the channel: 9
//=> chan consumer reads item from the channel: 10
//=> chan consumer reads item from the channel: 11
//=> Result: 166
//
//Example 3
//=> Message Hello 0
//=> Message Hello 1
//=> Message Hello 2
//=> Message Hello 3
//=> Message Hello 4
//=> Message Hello 5
//=> Message Hello 6
//=> Message Hello 7
//=> Message Hello 8
//=> Message Hello 9
//
//Example 4
//=> Message Bye 0
//=> Message Hello 0
//=> Message Hello 1
//=> Message Hello 2
//=> Message Bye 1
//=> Message Bye 2
//=> Message Hello 3
//=> Message Hello 4
//=> Message Bye 3
//=> Message Hello 5
//=> Message Hello 6
//=> Message Bye 4
//=> Message Bye 5
//=> Message Hello 7
//=> Message Bye 6
//=> Message Bye 7
//=> Message Hello 8
//=> Message Bye 8
//=> Message Hello 9
//=> Message Bye 9