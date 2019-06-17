package main

import (
    "fmt"
    "math/rand"
    "sync"
    "time"
)

const messageCount = 5

func main() {
    terminate := make(chan bool)

    fmt.Println("Example 1")

    c := work1(terminate)
    for i := 0; i < messageCount; i++{
        fmt.Printf("=> %s\n", <-c)
    }

    terminate <- true

    // If there are lots of goroutines,
    // this logis does not guarantee termination of all goroutines
    <-terminate

    fmt.Println("\nExample 2")

    var wg sync.WaitGroup
    wg.Add(2)

    c1 := work2(terminate, &wg, "goroutine-1")
    c2 := work2(terminate, &wg, "goroutine-2")

    go func() {
        for {
        select {
        case m := <-c1:
            fmt.Printf("=> %s\n", m)
        case m := <-c2:
            fmt.Printf("=> %s\n", m)
        case <- terminate:
            fmt.Println("=> Work has been done!")
            terminate <- true
            return
        default:
            fmt.Println("=> Waiting...")
            time.Sleep(500 * time.Millisecond)
        }
    } 

    }()

    wg.Wait()
    terminate <- true
    <- terminate
}

func work1(terminate chan bool) <-chan string {
    c := make(chan string)
    go func() {
        for {
            select {
            case c <- fmt.Sprintf("Message %d", rand.Intn(100)):
                time.Sleep(time.Duration(rand.Intn(300)) * time.Millisecond)
            case <- terminate:
                fmt.Println("=> Work has been done!")
                terminate <- true
                return
            }
        }
    }()
    return c
}

func work2(terminate chan bool, wg *sync.WaitGroup, name string) <-chan string {
    c := make(chan string)
    go func() {
        for i := 0; i < messageCount; i++ {
            c <- fmt.Sprintf("Message %s %d", name, i)
            time.Sleep(3 * time.Second)
        }
        wg.Done()
    }()
    return c
}

//Output
//Example 1
//=> Message 81
//=> Message 47
//=> Message 81
//=> Message 25
//=> Message 56
//=> Work has been done!
//
//Example 2
//=> Waiting...
//=> Message goroutine-1 0
//=> Message goroutine-2 0
//=> Waiting...
//=> Waiting...
//=> Waiting...
//=> Waiting...
//=> Waiting...
//=> Waiting...
//=> Message goroutine-2 1
//=> Message goroutine-1 1
//=> Waiting...
//=> Waiting...
//=> Waiting...
//=> Waiting...
//=> Waiting...
//=> Waiting...
//=> Message goroutine-1 2
//=> Message goroutine-2 2
//=> Waiting...
//=> Waiting...
//=> Waiting...
//=> Waiting...
//=> Waiting...
//=> Waiting...
//=> Message goroutine-2 3
//=> Message goroutine-1 3
//=> Waiting...
//=> Waiting...
//=> Waiting...
//=> Waiting...
//=> Waiting...
//=> Waiting...
//=> Message goroutine-1 4
//=> Message goroutine-2 4
//=> Waiting...
//=> Waiting...
//=> Waiting...
//=> Waiting...
//=> Waiting...
//=> Waiting...
//=> Work has been done!
