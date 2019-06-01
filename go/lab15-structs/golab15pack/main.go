package golab15pack

import "fmt"

// Structure name starts from lower case
// but if the 'public' method returns the structure
// all fields are visible at invocation method
type example1 struct {
    Flag bool
    counter int16
    pi float32
}

type example2 struct {
    flag bool
    counter int16
    pi float32
}

func M1() example1 {
    var e example1
    fmt.Printf("=> golab15pack.M1: %+v\n", e)
    //=> golab15pack.M1: {Flag:false counter:0 pi:0}
    return e
}

func M2() (e example1) {
    // It's possible to skip some field like in this section 
    e = example1 {
        Flag: true,
        pi: 3.14,
    }
    fmt.Printf("=> golab15pack.M2: %+v\n", e)
    //=> golab15pack.M2: {Flag:true counter:0 pi:3.14}
    return
}

func M3() (e example1) {
    // Structure with default values
    e = example1 {}
    fmt.Printf("=> golab15pack.M3: %+v\n", e)
    //=> golab15pack.M3: {Flag:false counter:0 pi:0}
    return
}

func M4() (e example1) {
    // Should be specified all fields in defined order
    e = example1 {true, 5, 3.14}
    fmt.Printf("=> golab15pack.M4: %+v\n", e)
    //=> golab15pack.M4: {Flag:true counter:5 pi:3.14}
    return
}

func M5() (e example2) {
    // Should be specified all fields in defined order
    e = example2 {true, 10, 3.14}
    fmt.Printf("=> golab15pack.M5: %+v\n", e)
    //=> golab15pack.M5: {Flag:true counter:10 pi:3.14}

    fmt.Println("=> golab15pack.M5 - e.flag:", e.flag)
    //=> golab15pack.M5 - e.flag: true

    fmt.Println("=> golab15pack.M5 - e.counter:", e.counter)
    //=> golab15pack.M5 - e.counter: 10

    fmt.Println("=> golab15pack.M5 - e.pi:", e.pi)
    //=> golab15pack.M5 - e.pi: 3.14

    return
}