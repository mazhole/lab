package main

import (
    "fmt"
    "math"
)

type Vertex struct {
    X, Y float64
}

func (v Vertex) Abs() float64 {
    return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

// Works with specified object
func (v *Vertex) Scale(f float64) {
    v.X = v.X * f
    v.Y = v.Y * f
}

// Works with the copy of specified object
func (v Vertex) noScale(f float64) {
    v.X = v.X * f
    v.Y = v.Y * f
    fmt.Println("=> v.noScale:", v.Abs())
    //=> v.noScale: 500
}

func main() {
    v := Vertex{3, 4}
    fmt.Println("=> v:", v.Abs())
    //=> v: 5

    v.Scale(10)
    fmt.Println("=> v:", v.Abs())
    //=> v: 50

    v.noScale(10)
    fmt.Println("=> v has not been changed:", v.Abs())
    //=> v has not been changed: 50
}
