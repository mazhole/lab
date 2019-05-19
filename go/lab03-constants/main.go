package main

const someInt = 1
const typedInt int32 = 17
const str = "string value"

const (
    c1 = 1
    c2 = 2
)

//like enums
const (
    one = iota + c1 //1
    two             //2
    _               //skip iota
    four            //4
)

const (
    _ = iota //skipping the first value
    x1 = 2 * iota + 1
    x2
    x3
    x4
)

func main() {
    var pi float32 = 3.14

    println(pi + someInt)
    //+4.140000e+000

    println(pi + float32(typedInt))
    //+2.014000e+001

    println(c1, c2)
    //1 2

    println(one, two, four)
    //1 2 4

    println(x1, x2, x3, x4)
    //3 5 7 9
}
