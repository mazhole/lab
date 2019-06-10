package main

import "fmt"

func main() {
    do("work in bounds of 'main' goroutine")

    go do("work in bounds of another goroutine")

    for i := 0; i < 20; i++ {
        go do(fmt.Sprintf("work in bounds of another goroutine: %d", i))

        go func(){
            do("work in bounds of anonymous function and another goroutine")
        }()
    }

    fmt.Scanln()
}

func do(action string) {
    fmt.Println("=> do action:", action)
}

//Output:
//=> do action: work in bounds of 'main' goroutine
//=> do action: work in bounds of anonymous function and another goroutine
//=> do action: work in bounds of anonymous function and another goroutine
//=> do action: work in bounds of another goroutine: 1
//=> do action: work in bounds of anonymous function and another goroutine
//=> do action: work in bounds of another goroutine: 2
//=> do action: work in bounds of another goroutine: 8
//=> do action: work in bounds of another goroutine: 0
//=> do action: work in bounds of anonymous function and another goroutine
//=> do action: work in bounds of another goroutine: 3
//=> do action: work in bounds of anonymous function and another goroutine
//=> do action: work in bounds of another goroutine: 4
//=> do action: work in bounds of anonymous function and another goroutine
//=> do action: work in bounds of another goroutine: 5
//=> do action: work in bounds of anonymous function and another goroutine
//=> do action: work in bounds of another goroutine: 6
//=> do action: work in bounds of anonymous function and another goroutine
//=> do action: work in bounds of anonymous function and another goroutine
//=> do action: work in bounds of another goroutine: 9
//=> do action: work in bounds of anonymous function and another goroutine
//=> do action: work in bounds of another goroutine: 10
//=> do action: work in bounds of anonymous function and another goroutine
//=> do action: work in bounds of another goroutine: 11
//=> do action: work in bounds of anonymous function and another goroutine
//=> do action: work in bounds of another goroutine: 12
//=> do action: work in bounds of anonymous function and another goroutine
//=> do action: work in bounds of another goroutine: 13
//=> do action: work in bounds of anonymous function and another goroutine
//=> do action: work in bounds of anonymous function and another goroutine
//=> do action: work in bounds of another goroutine: 14
//=> do action: work in bounds of another goroutine: 7
//=> do action: work in bounds of anonymous function and another goroutine
//=> do action: work in bounds of another goroutine: 15
//=> do action: work in bounds of anonymous function and another goroutine
//=> do action: work in bounds of another goroutine: 16
//=> do action: work in bounds of another goroutine: 17
//=> do action: work in bounds of anonymous function and another goroutine
//=> do action: work in bounds of another goroutine: 19
//=> do action: work in bounds of anonymous function and another goroutine
//=> do action: work in bounds of anonymous function and another goroutine
//=> do action: work in bounds of another goroutine: 18
//=> do action: work in bounds of another goroutine