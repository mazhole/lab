package main

import "fmt"

type iPrintInfo interface {
    Print()
}

type intInfo int

func (ii intInfo) Print() {
    fmt.Println("=> info: ", ii)
}

type stringInfo string

func (si stringInfo) Print() {
    fmt.Println("=> info: ", si)
}

type info struct {
    id int
    iPrintInfo
}

func (i info) PrintInfo() {
    fmt.Println("=> id:", i.id)
    i.Print()
}

func main() {
    ii1 := intInfo(5)
    si1 := stringInfo("Some data")

    i1 := info {1, ii1}
    i1.PrintInfo()
    //=> id: 1
    //=> info:  5

    i2 := info {2, si1}
    i2.PrintInfo()
    //=> id: 2
    //=> info:  Some data
}