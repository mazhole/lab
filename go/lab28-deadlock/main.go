package main

import (
    "fmt"
    "time"
)

const hitsCount = 5

type ball struct {
    hits int
}


func main() {
    table := make(chan *ball)

    // Uncomment to execute with deadlock!!!
    //table <- new(ball)

    go play("ping", table)
    go play("pong", table)

    // Uncomment to execute without deadlock!!!
    //table <- new(ball)
    time.Sleep(15 * time.Second)

    fmt.Println("=> Take the ball")
    <-table

    fmt.Scanln()
}

func play(name string, table chan *ball) {
    for i := 0; i < hitsCount; i++ {
        ball := <-table
        ball.hits++
        fmt.Printf("=> %s: hits %d\n", name, ball.hits)
        time.Sleep(time.Second)
        table <- ball
    }
}

//Output

//=> Take the ball
//fatal error: all goroutines are asleep - deadlock!

//or

//=> pong: hits 1
//=> ping: hits 2
//=> pong: hits 3
//=> ping: hits 4
//=> pong: hits 5
//=> ping: hits 6
//=> pong: hits 7
//=> ping: hits 8
//=> pong: hits 9
//=> ping: hits 10
//=> Take the ball