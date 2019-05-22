package main

// Package fmt implements formatted I/O with functions analogous to C's
// printf and scanf. The format 'verbs' are derived from C's but are simpler.
import "fmt"

const size uint = 3

func main() {
    var m1 map[string]string
    fmt.Println("=> uninitialized map m1:", m1)
    //=> uninitialized map m1: map[]

    // runtime error: assignment to entry in nil map
    //m1["0"] = "one"

    // long map declaration and initialization
    //var m2 map[string]string = map[string]struct{}
    
    // short map declaration and initialization
    m2 := map[string]string{}
    m2["0"] = "zero"
    fmt.Println("=> initialized map m2:", m2)
    //=> initialized map m2: map[0:zero]

    var m3 = make(map[string]string)
    m3["00"] = "zero-zero"
    fmt.Println("=> initialized map m3:", m3)
    //=> initialized map m3: map[00:zero-zero]

    val1 := m3["00"]
    fmt.Println("=> first item value of m3:", val1)
    //=> first item value of m3 zero-zero

    val1 = "one-one"
    fmt.Println("=> initialized map m3:", m3)
    //=> initialized map m3: map[00:zero-zero]

    val2 := m3["11"]
    fmt.Println("=> unexisted item value of m3 is empty string:", val2, "length val2:", len(val2))
    //=> unexisted item value of m3 is empty string:  length val2: 0

    // 'val3' contains item value, 'isExists' contains existence info
    val3, isExists := m3["11"]
    fmt.Println("=> val3:", val3, "isExists:", isExists)
    //=> val3:  isExists: false

    // Gets only existence into
    _, exists := m3["11"]
    fmt.Println("=> exists:", exists)
    //=> exists: false

    key := "00"
    delete(m3, key)
    _, exists = m3["00"]
    fmt.Println("=> exists:", exists)
    //=> exists: false

    // m4 contains the reference to the map
    m4 := m3
    m4[key] = "test"
    fmt.Println("=> map m3:", m3)
    fmt.Println("=> map m4:", m4)
    //=> map m3: map[00:test]
    //=> map m4: map[00:test]
}
