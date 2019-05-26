package main

import "fmt"

func main() {
    a := 1
    b := &a
    fmt.Println("=> a:", a, "b:", *b)
    //=> a: 2 b: 1

    *b = 2
    fmt.Println("=> a:", a, "b:", *b)
    //=> a: 2 b: 2

    //slice
    s := []int{0, 0, 0, 0}
    fmt.Println("=> slice:", s)
    //=> slice: [0 0 0 0]

    updateSlice(s)
    fmt.Println("=> slice:", s)
    //=> slice: [1 0 0 0]
}

func updateSlice(s []int) {
    s[0] = 1
}