package main

import (
    "errors"
    "fmt"
)

//type error interface {
//    error() string
//}

type errorString1 string

func (e errorString1) Error() string {
    return string(e)
}

func New1(text string) error {
    return errorString1(text)
}

// Like in https://github.com/golang/go/blob/master/src/errors/errors.go
type errorString2 struct {
    s string
}

// Like in https://github.com/golang/go/blob/master/src/errors/errors.go
func New2(text string) error {
    return &errorString2{text}
}

// Like in https://github.com/golang/go/blob/master/src/errors/errors.go
func (e *errorString2) Error() string {
    return e.s
}

const errorInfo = "Error's info"

var Error1 = New1(errorInfo)
var Error2 = New2(errorInfo)
var Error3 = errors.New(errorInfo)

type LabError struct {
    Line int
    File string
    Message string
}

func (e *LabError) Error() string {
    return fmt.Sprintf("=> %s:%d: %s", e.File, e.Line, e.Message)
}

type temporary interface {
    Temporary() bool
}

func main() {
    fmt.Printf("=> Type: %T Value: %+v\n", Error1, Error1)
    //=> Type: main.errorString1 Value: Error's info

    // True...
    if Error1 == New1(errorInfo) {
        fmt.Println("=> Error1")
        //=> Error1
    }

    fmt.Printf("=> Type: %T Value: %+v\n", Error2, Error2)
    //=> Type: *main.errorString2 Value: Error's info

    // False... (different references)
    if Error2 == New2(errorInfo) {
        fmt.Println("=> Error2")
    }

    fmt.Printf("=> Type: %T Value: %+v\n", Error3, Error3)
    //=> Type: *errors.errorString Value: Error's info

    // False... (different references)
    if Error3 == errors.New(errorInfo) {
        fmt.Println("=> Struct Type Error")
    }

    err := BadFunc()

    // Switch type assertion
    switch err := err.(type) {
    case nil:
        fmt.Println("=> Everything is OK!")
    case *LabError:
        fmt.Printf("=> Error!!! [Type]: %T [Value]: %+v [Line]: %v\n", err, err, err.Line)
    default:
        fmt.Printf("=> Unknown error. Type: %T Value: %+v\n", err, err)
    }
    //=> Error!!! [Type]: *main.LabError [Value]: => Something goes wrong:123: main.go [Line]: 123
}

func BadFunc() error {
    return &LabError{123, "Something goes wrong", "main.go"}
}
