package main

import (
    "fmt"
    "time"
)

func main() {
    showMeTheMoney()
    //$$$$

    s := sum(1, 2)
    fmt.Println("=> sum1:", s)
    //=> sum1: 3

    s = sum(2, 3)
    fmt.Println("=> sum2:", s)
    //=> sum2: 5

    stuff := []int{10, 20, 30}
    s = sumMore(stuff...)
    fmt.Println("=> stuff sum:", s)
    //=> stuff sum: 6

    res1, err1 := sumOnlyNatural1(stuff...)
    fmt.Println("=> stuff:", stuff)
    //=> stuff: [10 20 30]
    fmt.Println("=> res1:", res1, "err1:", err1)
    //=> res1: 60 err1: <nil>

    stuff = append(stuff, -1)
    res1, err1 = sumOnlyNatural1(stuff...)
    fmt.Println("=> stuff:", stuff)
    //=> stuff: [10 20 30 -1]
    fmt.Println("=> res1:", res1, "err1:", err1)
    //=> res1: 0 err1: Only natural numbers expected: -1 less than 0

    res1, err1 = sumOnlyNatural2(stuff...)
    fmt.Println("=> stuff:", stuff)
    //=> stuff: [10 20 30 -1]
    fmt.Println("=> res1:", res1, "err1:", err1)
    //=> res1: 0 err1: Only natural numbers expected: -1 less than 0

    // https://github.com/golang/go/wiki/SliceTricks
    //stuff = append(stuff[:len(stuff) -1], []int{}...)
    stuff = stuff[:len(stuff) - 1]
    res1, err1 = sumOnlyNatural2(stuff...)
    fmt.Println("=> stuff:", stuff)
    //=> stuff: [10 20 30]
    fmt.Println("=> res1:", res1, "err1:", err1)
    //=> res1: 60 err1: <nil>

    myTimer := getTimer()
    //https://habr.com/ru/post/118898/
    defer myTimer()

    f := func() {
        showMeTheMoney()
    }

    time.Sleep(1 * time.Second)

    f()
    //$$$$
}
//Defer:
//=> Time from start 1.000244168s

func showMeTheMoney() {
    fmt.Println("$$$$")
}

func sum(i int, j int) int {
    return i + j
}

func sumLight(i, j int) int {
    return i + j
}

// Returns naned data 'res'
func sumMore(stuff ...int) (res int) {
    for _, v := range stuff {
        res += v
    }
    return
}

func sumOnlyNatural1(stuff ...int) (int, error) {
    res := 0
    for _, v := range stuff {
        if v < 0 {
            return 0, fmt.Errorf("Only natural numbers expected: %d less than 0", v)
        }
        res += v
    }

    return res, nil
}

// Returns named data 'res' and 'error'
func sumOnlyNatural2(stuff ...int) (res int, err error) {
    for _, v := range stuff {
        if v < 0 {
            err = fmt.Errorf("Only natural numbers expected: %d less than 0", v)

            //return 0
            //compilation error:
            //not enough arguments to return
            //    have (number)
            //    want (int, error)

            return 0, err
        }
        res += v
    }

    return
}

func getTimer() func() {
    start := time.Now()
    return func() {
        fmt.Printf("=> Time from start %v", time.Since(start))
    }
}