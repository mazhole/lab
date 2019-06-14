package main

import (
    "fmt"
    "math/rand"
    //"runtime"
    "time"
)

// Operator 'select' allows to implement non-blocking operations.
// It can combine sending/receiving to/from different channels.

// If there is no data in the channels and default
// section is not specified, goroutine is blocked.

// If two or more channels contain data, runtime
// gets data from some of the channel.

const messageCount = 5


func main() {

    // https://golang.org/pkg/runtime/#GOMAXPROCS
    //runtime.GOMAXPROCS(runtime.NumCPU())


    fmt.Println("Example 1")

    c1 := make(chan string)
    c2 := make(chan string)

    go func() {
        for i := 0; i < messageCount; i++ {
            c1 <- fmt.Sprintf("[channel 1] info %d", i)
            time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
        }
    }()

    go func() {
        for i := 0; i < messageCount; i++ {
            c2 <- fmt.Sprintf("[channel 2] info %d", i)
            time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
        }
    }()

    // Here message cannot be sent to the messages channel,
    // because the channel has no buffer and there is no receiver.
    // Therefore the default case is selected.
    select {
    case c1 <- "[channel 1] additional message":
        fmt.Println("=> Additional message is sent to [channel 1]")
    case c2 <- "[channel 2] additional message":
        fmt.Println("=> Additional message is sent to [channel 2]")
    default:
        time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
        fmt.Println("=> No message is sent")
    }

    go func() {
        for i := 0; i < messageCount * 3; i++ {
            select {
            case msg := <- c1:
                fmt.Println("=>", msg)
            case msg := <- c2:
                fmt.Println("=>", msg)
            default:
                fmt.Println("=> Default message")
                time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
            }

            // https://golang.org/pkg/runtime/#Gosched
            //runtime.Gosched()
        }
    }()

    time.Sleep(time.Duration(20) * time.Second)

    fmt.Println("\nExample 2")

    c3 := channelProvider("Hello")
    c4 := channelProvider("Bye")
    c5 := multiChannel(c3, c4)
    for i := 0; i < messageCount * 2; i++ {
        fmt.Println("=>", <-c5)
    }

    fmt.Scanln()
}

func multiChannel(c1, c2 <- chan string) <-chan string {
    c := make(chan string)
    go func() {
        for {
            select {
            case msg := <- c1:
                c <- msg
            case msg := <- c2:
                c <- msg
            }
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
//=> No message is sent
//=> [channel 2] info 0
//=> [channel 1] info 0
//=> Default message
//=> [channel 1] info 1
//=> Default message
//=> [channel 1] info 2
//=> [channel 2] info 1
//=> Default message
//=> Default message
//=> Default message
//=> [channel 1] info 3
//=> Default message
//=> Default message
//=> [channel 1] info 4
//=> [channel 2] info 2
//
//Example 2
//=> Message Bye 0
//=> Message Hello 0
//=> Message Hello 1
//=> Message Bye 1
//=> Message Hello 2
//=> Message Bye 2
//=> Message Bye 3
//=> Message Hello 3
//=> Message Bye 4
//=> Message Hello 4