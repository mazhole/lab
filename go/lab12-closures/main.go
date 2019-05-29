package main

import (
    "fmt"
    "time"
)

func main() {
    myTimer := getTimer()
    myTimer()
    //=> Time from start: 131.133µs
}

func getTimer() func() {
    start := time.Now()
    fmt.Println("=> start:", start)
    //=> start: 2019-05-13 20:19:37.46084684 +0300 MSK m=+0.000032616
    return func() {
        fmt.Printf("=> Time from start: %v", time.Since(start))
    }
}