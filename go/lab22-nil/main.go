package main

import "fmt"

type myError string

func (m myError) Error() string {
    return string(m)
}

func main() {
    var i interface{}

    fmt.Printf("=> %v %v\n", i, i == nil)
    //=> <nil> true

    s := []int{1, 2, 3}
    i = s

    fmt.Printf("=> %v %v\n", i, i == nil)
    //=> [1 2 3] false

    if f := getError1(10); f != nil {
        fmt.Println("=> f is not nill")
        //=> f is not nill
    }

    if f := getError2(); f == nil {
        fmt.Println("=> f is nil")
        //=> f is nil
    }
}

func getError1(input interface{}) error {
    var m *myError
    if _, ok := input.(int); ok {
        fmt.Println("=> getError1: input is int")
        //=> getError1: input is int
        return m
    }
    return nil
}

func getError2() error {
    return nil
}
