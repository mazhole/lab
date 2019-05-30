package main

import (
    "fmt"
    "io"
    "os"
)

//https://habr.com/ru/post/118898/

func main() {
    DeferExample();

    res := DeferAndClosureExample()
    fmt.Println("=> DeferAndClosureExample returns", res)
    //=> DeferAndClosureExample returns 2

    CopyFile("run1.sh", "run.sh")
}

func DeferExample() {
    fmt.Println("=> DeferExample begin")
    //=> DeferExample begin

    for i := 0; i < 5; i++ {
        defer fmt.Println("=> i:", i)
    }

    fmt.Println("=> DeferExample end")
    //=> DeferExample end
}
//=> i: 4
//=> i: 3
//=> i: 2
//=> i: 1
//=> i: 0

func DeferAndClosureExample() (i int) {
    fmt.Println("=> DeferAndClosureExample begin")
    //=> DeferAndClosureExample begin

    //It there are lots of methods in the call chain, only the last method (ex. m3) is placed into defer.
    //Previous methods (ex. m1 and m2) are executed in bounds of covered method.
    //defer o.m1().m2().m3()

    defer func() {
        i++
    }()

    fmt.Println("=> DeferAndClosureExample end")
    //=> DeferAndClosureExample end

    return 1
}

func CopyFile(dstName, srcName string) (written int64, err error) {
    src, err := os.Open(srcName)//, os.O_RDONLY, 0)
    fmt.Println("=> Read source file", srcName)
    //=> Read source file run.sh
    if err != nil {
        return
    }

    dst, err := os.Create(dstName)//, os.O_WRONLY|os.O_CREATE, 0644)
    fmt.Println("=> Write destination file", dstName)
    //=> Write destination file run1.sh
    if err != nil {
        return
    }

    written, err = io.Copy(dst, src)
    dst.Close()
    src.Close()
    return
}