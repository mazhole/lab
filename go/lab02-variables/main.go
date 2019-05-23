package main

import (
    "bytes"
)

//package variables
var (
    p0 = "Hello"
    p1 = "World"
)

func main() {
    var i int = 10 //depends on the platform: 32 bit or 64 bit
    var autoInt = -10
    var bigInt int64 = 1<<32 - 1
    var unsignedInt uint = 100500 //depends on the platform: 32 bit or 64 bit
    var unsignedBigInt uint64 = 1<<64 - 1
    println("=> integers:", i, autoInt, bigInt, unsignedInt, unsignedBigInt)
    //=> integers: 10 -10 4294967295 100500 18446744073709551615

    var p float32 = 3.14 //float: float32, float64
    println("=> float:", p)
    //=> float: +3.140000e+000

    var b = true
    println("=> bool:", b)
    //=> bool: true

    var c, d int = 1, 2
    println("=> c:", c)
    //=> c: 1
    println("=> d:", d)
    //=> d: 2

    var hello string = "Hello\n\t\t"
    var world = "World!"
    println("=> string:", hello, world)
    //=> string: Hello
    //         World!

    var binary byte = '\x27'
    println("=> binary:", binary)
    //=> binary: 39

    //short defining
    meaningOfLife := 42
    println("=> Meaning of life:", meaningOfLife)
    //=> Meaning of life: 42

    println("=> float to int convertion:", int(p))
    //=> float to int convertion: 3
    println("=> int to string convertion:", string(48))
    //=> int to string convertion: 0

    //complex numbers
    z := 2 + 3i
    println("=> complex number:", z)
    //=> complex number: (+2.000000e+000+3.000000e+000i)

    //concatenation of strings
    s1 := "s1 "
    s2 := "s2"
    concat := s1 + s2
    println("=> concatenation of strings:", concat, len(concat))
    //=> concatenation of strings: s1 s2 5

    s3 := `Hello\r\n
    World`
    println("=> string as-is:", s3)
    //=> string as-is: Hello\r\n
    //    World

    var defaultInt int
    var defaultFloat float32
    var defaultString string
    var defaultBool bool
    println("=> default values:", defaultInt, defaultFloat, defaultString, defaultBool)
    //=> default values: 0 +0.000000e+000  false

    //defining of variables
    var v1, v2 string = "v1", "v2"
    println(v1, v2)
    //v1 v2

    var (
        m0 int = 10
        m1 = "str"
        m2 = 23
    )
    println(m0, m1, m2)
    //10 str 23

    println(p0, p1)
    //Hello World

    var buffer bytes.Buffer

    for i := 0; i < 100; i++ {
        buffer.WriteString("a")
    }

    println("=> concatenated string:", buffer.String())
    //=> concatenated string: aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa
}
