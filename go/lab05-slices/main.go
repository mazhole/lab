package main

// Package fmt implements formatted I/O with functions analogous to C's
// printf and scanf. The format 'verbs' are derived from C's but are simpler.
import "fmt"

const size uint = 3

func main() {
    var s1 []int

    fmt.Println("=> empty slice s1:", s1, "length:", len(s1), "capacity:", cap(s1))
    //=> empty slice s1: [] length: 0 capacity: 0

    // runtime error: index out of range
    //fmt.Println("=> item of slice s1:", s1[5])

    s1 = append(s1, 100)
    fmt.Println("=> slice s1:", s1, "length:", len(s1), "capacity:", cap(s1))
    //=> slice s1: [100] length: 1 capacity: 1
    s1 = append(s1, 101)
    fmt.Println("=> slice s1:", s1, "length:", len(s1), "capacity:", cap(s1))
    //=> slice s1: [100 101] length: 2 capacity: 2
    s1 = append(s1, 102)
    fmt.Println("=> slice s1:", s1, "length:", len(s1), "capacity:", cap(s1))
    //=> slice s1: [100 101 102] length: 3 capacity: 4
    s1 = append(s1, 103)
    fmt.Println("=> slice s1:", s1, "length:", len(s1), "capacity:", cap(s1))
    //=> slice s1: [100 101 102 103] length: 4 capacity: 4
    s1 = append(s1, 104)
    fmt.Println("=> slice s1:", s1, "length:", len(s1), "capacity:", cap(s1))
    //=> slice s1: [100 101 102 103 104] length: 5 capacity: 8

    // ... => initialized length
    s2 := []int{1, 2, 3}
    fmt.Println("=> slice s2:", s2, "length:", len(s2), "capacity:", cap(s2))
    //=> slice s2: [1 2 3] length: 3 capacity: 3

    // Adds slice to another slice by elements
    s1 = append(s1, s2...)
    fmt.Println("=> slice s1:", s1, "length:", len(s1), "capacity:", cap(s1))
    //=> slice s1: [100 101 102 103 104 1 2 3] length: 8 capacity: 8

    var s3 [][]int
    s3 = append(s3, s1, s2)
    fmt.Println("=> slice s3:", s3, "length:", len(s3), "capacity:", cap(s3))
    //=> slice s3: [[100 101 102 103 104 1 2 3] [1 2 3]] length: 2 capacity: 2

    // Creates slice with default values and specified length
    s4 := make([]int, 10)
    fmt.Println("=> slice s4:", s4, "length:", len(s4), "capacity:", cap(s4))
    //=> slice s4: [0 0 0 0 0 0 0 0 0 0] length: 10 capacity: 10

    s4 = append(s4, 1)
    fmt.Println("=> slice s4:", s4, "length:", len(s4), "capacity:", cap(s4))
    //=> slice s4: [0 0 0 0 0 0 0 0 0 0 1] length: 11 capacity: 20
    
    // Creates slice with default values and specified length and capacity
    s5 := make([]int, 10, 15)
    fmt.Println("=> slice s5:", s5, "length:", len(s5), "capacity:", cap(s5))
    //=> slice s5: [0 0 0 0 0 0 0 0 0 0] length: 10 capacity: 15
    
    s5 = append(s5, 1)
    fmt.Println("=> slice s5:", s5, "length:", len(s5), "capacity:", cap(s5))
    //=> slice s5: [0 0 0 0 0 0 0 0 0 0 1] length: 11 capacity: 15

    s5 = append(s5, []int{1, 2, 3, 4, 5}...)
    fmt.Println("=> slice s5:", s5, "length:", len(s5), "capacity:", cap(s5))
    //=> slice s5: [0 0 0 0 0 0 0 0 0 0 1 1 2 3 4 5] length: 16 capacity: 30

    // Slice variable contains reference to the slice
    s6 := s5
    s6[0] = 42
    fmt.Println("=> slice s5:", s5, "length:", len(s5), "capacity:", cap(s5))
    //=> slice s5: [42 0 0 0 0 0 0 0 0 0 1 1 2 3 4 5] length: 16 capacity: 30
    fmt.Println("=> slice s6:", s6, "length:", len(s6), "capacity:", cap(s6))
    //=> slice s6: [42 0 0 0 0 0 0 0 0 0 1 1 2 3 4 5] length: 16 capacity: 30

    s5 = append(s5, []int{1, 2, 3, 4, 5}...)
    fmt.Println("=> slice s5:", s5, "length:", len(s5), "capacity:", cap(s5))
    //=> slice s5: [42 0 0 0 0 0 0 0 0 0 1 1 2 3 4 5 1 2 3 4 5] length: 21 capacity: 30
    fmt.Println("=> slice s6:", s6, "length:", len(s6), "capacity:", cap(s6))
    //=> slice s6: [42 0 0 0 0 0 0 0 0 0 1 1 2 3 4 5] length: 16 capacity: 30

    // Wrong copy. Function 'copy' copies slice on the base of capacity of target slice
    var s7 []int
    copy(s7, s6)
    fmt.Println("=> slice s7:", s7, "length:", len(s7), "capacity:", cap(s7))
    //=> slice s7: [] length: 0 capacity: 0

    s8 := make([]int, 2, 2)
    copy(s8, s6)
    fmt.Println("=> slice s8:", s8, "length:", len(s8), "capacity:", cap(s8))
    //=> slice s8: [42 0] length: 2 capacity: 2

    s9 := make([]int, len(s6), cap(s6))
    copy(s9, s6)
    fmt.Println("=> slice s9:", s9, "length:", len(s9), "capacity:", cap(s9))
    //=> slice s9: [42 0 0 0 0 0 0 0 0 0 1 1 2 3 4 5] length: 16 capacity: 30

    // Slice part: [...)
    fmt.Println("=> slice s9 parts:", s9[:12], s9[10:12], s9[12:])
    //=> slice s9 parts: [42 0 0 0 0 0 0 0 0 0 1 1] [1 1] [2 3 4 5]
    s9 = append(s9[:12], s9[14:]...)
    fmt.Println("=> slice s9:", s9, "length:", len(s9), "capacity:", cap(s9))
    //=> slice s9: [42 0 0 0 0 0 0 0 0 0 1 1 4 5] length: 14 capacity: 30

    // Slice on the base of array
    a1 := [...]int{3, 4, 5, 6}
    s10 := a1[:]
    a1[2] = 9
    fmt.Println("=> array a1:", a1, "length:", len(a1))
    //=> array a1: [3 4 9 6] length: 4
    fmt.Println("=> slice s10:", s10, "length:", len(s10), "capacity:", cap(s10))
    //=> slice s10: [3 4 9 6] length: 4 capacity: 4

    a2 := s10[:]
    fmt.Println("=> array a2:", a2, "length:", len(a2))
    //=> array a2: [3 4 9 6] length: 4
    a2[1] = 42
    fmt.Println("=> array a1:", a1, "length:", len(a1))
    //=> array a1: [3 42 9 6] length: 4
    fmt.Println("=> array a2:", a2, "length:", len(a2))
    //=> array a2: [3 42 9 6] length: 4
    fmt.Println("=> slice s10:", s10, "length:", len(s10), "capacity:", cap(s10))
    //=> slice s10: [3 42 9 6] length: 4 capacity: 4
}
