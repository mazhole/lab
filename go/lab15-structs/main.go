package main

import (
    "fmt"
    "github.com/mazhole/study/go/lab15-structs/golab15pack"
)

func main() {
    e1 := golab15pack.M1()
    fmt.Printf("=> main - golab15pack.M1: %+v\n", e1)
    //=> main - golab15pack.M1: {Flag:false counter:0 pi:0}

    e1 = golab15pack.M2()
    fmt.Printf("=> main - golab15pack.M1: %+v\n", e1)
    //=> main - golab15pack.M1: {Flag:true counter:0 pi:3.14}

    e1 = golab15pack.M3()
    fmt.Printf("=> main - golab15pack.M1: %+v\n", e1)
    //=> main - golab15pack.M1: {Flag:false counter:0 pi:0}

    e1 = golab15pack.M4()
    fmt.Printf("=> main - golab15pack.M1: %+v\n", e1)
    //=> main - golab15pack.M1: {Flag:true counter:5 pi:3.14}

    e2 := golab15pack.M5()
    fmt.Printf("=> main - golab15pack.M1: %+v\n", e2)
    //=> main - golab15pack.M1: {flag:true counter:10 pi:3.14}


    fmt.Println("=> main - e1.Flag:", e1.Flag)
    //=> main - e1.Flag: true

    //fmt.Println("=> main - e1.counter:", e1.counter)
    //compilation error: e1.counter undefined (cannot refer to unexported field or method counter)

    //fmt.Println("=> main - e1.pi:", e1.pi)
    //compilation error: e1.pi undefined (cannot refer to unexported field or method pi)


    //fmt.Println("=> main - e2.flag:", e2.flag)
    //compilation error: => e2.flag undefined (cannot refer to unexported field or method flag)

    //fmt.Println("=> main - e2.counter:", e2.counter)
    //compilation error: e2.counter undefined (cannot refer to unexported field or method counter)

    //fmt.Println("=> main - e2.pi:", e2.pi)
    //compilation error: e2.pi undefined (cannot refer to unexported field or method pi)
}
