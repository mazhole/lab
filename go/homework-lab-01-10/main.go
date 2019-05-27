// This code should be implemented by yourself!!!

package main

import (
    "fmt"
    "math/rand"
    "sort"
    "strings"
    "github.com/spf13/cast"
)

func main() {

}

func ReturnInt() int {
    return rand.Intn(10)
}

func ReturnFloat() float32 {
    return float32(2.2 / 2)
}

func ReturnIntArray() [3]int {
    a := [3]int{1, 3, 4}
    return a
}

func ReturnIntSlice() []int {
    //s := []int{2, 2, 3}
    s := []int{1, 2, 3}
    return s
}

func IntSliceToString(s []int) string {
    return strings.Trim(strings.Join(strings.Split(fmt.Sprint(s), " "), ""), "[]")
}

func MergeSlices1(s1 []float32, s2 []int32) []int {
    var result []int

    for _, v := range s1 {
        result = append(result, int(v))
    }

    //fmt.Println(cast.ToIntSlice(interface{}(s2)))
    result = append(result, cast.ToIntSlice(interface{}(s2))...) 
    
    return result
}

func MergeSlices2(s1 []float32, s2 []int32) []int32 {
    var result []int32

    for _, v := range s1 {
        result = append(result, int32(v))
    }

    for _, v := range s2 {
        result = append(result, int32(v))
    } 
    
    return result
}

//https://golang.org/pkg/sort/
func GetMapValuesSortedByKey(m map[int]string) []string {
    keys := make([]int, 0, len(m))
    for k := range m {
        keys = append(keys, k)
    }

    sort.Ints(keys)

    //fmt.Println(keys)

    values := make([]string, 0, len(m))
    for _, k := range keys {
        //fmt.Println(k)
        values = append(values, m[k])
    }

    return values
}