package main

// Package fmt implements formatted I/O with functions analogous to C's
// printf and scanf. The format 'verbs' are derived from C's but are simpler.
import "fmt"

const size uint = 3

func main() {
    var a1 [3]int

    fmt.Println("=> array a1:", a1, "length:", len(a1))
    //=> array a1: [0 0 0] length: 3

    var a2 [2 * size] bool

    fmt.Println("=> array a2:", a2, "length:", len(a2))
    //=> array a2: [false false false false false false] length: 6

    // ... => initialized length
    a3 := [...]int{1, 2, 3}
    fmt.Println("=> array a3:", a3, "length:", len(a3))
    //=> array a3: [1 2 3] length: 3

    a3[1] = 12
    fmt.Println("=> array a3:", a3, "length:", len(a3))
    //=> array a3: [1 12 3] length: 3

    // Compilaton error: invalid array index 7 (out of bounds for 3-element array)
    //a3[7] = 12

    var a4 [3][3]int
    a4[1][2] = 7
    fmt.Println("=> matrix a4:", a4)
    //=> matrix a4: [[0 0 0] [0 0 7] [0 0 0]]
}
