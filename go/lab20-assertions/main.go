package main

import "fmt"

// Type assertions
func main() {
    var i interface{} = "Hi!"

    s := i.(string)
    fmt.Println("=> ", s)
    //=>  Hi!

    s, ok := i.(string)
    fmt.Println("=> ", s, ok)
    //=>  Hi! true

    f, ok := i.(float64)
    fmt.Println("=> ", f, ok)
    //=>  0 false

    //f = i.(float64) // panic: interface conversion: interface {} is string, not float64
    //fmt.Println(f)

    slice1 := []int{10, 12, 15}
    var sliceInterface interface{}
    sliceInterface = slice1

    sliceOfInterfaces := DoInterfaces(sliceInterface)

    fmt.Printf("=> slice1. Type: %T Value: %+v\n", slice1, slice1)
    //=> slice1. Type: []int Value: [10 12 15]

    fmt.Printf("=> sliceInterface. Type: %T Value: %+v\n", sliceInterface, sliceInterface)
    //=> sliceInterface. Type: []int Value: [10 12 15]

    fmt.Printf("=> sliceOfInterfaces. Type: %T Value: %+v\n", sliceOfInterfaces, sliceOfInterfaces)
    //=> sliceOfInterfaces. Type: []interface {} Value: [10 12 15]
}

func DoInterfaces(sliceInterface interface{}) (sliceOfInterfaces []interface{}) {
    s := sliceInterface.([]int)
    sliceOfInterfaces = make([]interface{}, len(s))
    for i, v := range s {
        sliceOfInterfaces[i] = v
    }
    return
}