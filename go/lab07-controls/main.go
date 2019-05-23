package main

const size uint = 3

func main() {
    isTrue := true
    if isTrue {
        println("=> isTrue:", isTrue)
        //=> isTrue: true
    }

    b := 1
    c := 0
    // b should be obligatorily checked on 1 value
    if b == 1 && isTrue && c != 1 {
        println("=> b should be obligatorily checked on 1 value")
        //=> b should be obligatorily checked on 1 value
    }

    m := map[string]string {
        "0": "zero",
        "1": "one",
    }
    if zeroValue, exists := m["0"]; exists {
        println("=> 0 item exists", zeroValue)
    } else {
        println("=> 0 item doesn't exist")
    }
    //=> 0 item exists zero

    for {
        println("=> endless loop")
        //=> endless loop
        break
    }

    s := []int{10, 11, 12, 13, 14, 15, 16, 17, 18, 19}
    v := 0;

    for i := 0; i < 4; {
        if i < 2 {
            i++
            continue
        }
        v = s[i]
        i++
        println("=> v:", v, "i:", i)
        //=> v: 2 i: 3
        //=> v: 3 i: 4
    }

    for i := 0; i < len(s); i++ {
        println("=> c-style loop", i, s[i])
        //=> c-style loop 0 10
        //=> c-style loop 1 11
        //...
        //=> c-style loop 9 19
    }

    for i := range s {
        println("=> range slice by index", i, s[i])
        //=> range slice by index 0 10
        //=> range slice by index 1 11
        //...
        //=> range slice by index 9 19
    }

    // Better range version
    for i, v := range s {
        println("=> range slice by index", i, v)
        //=> range slice by index 0 10
        //=> range slice by index 1 11
        //...
        //=> range slice by index 9 19
    }

    // ... and another variant
    for _, v := range s {
        println("=> range slice by index", v)
        //=> range slice by index 10
        //=> range slice by index 11
        //...
        //=> range slice by index 19
    }

    for k := range m {
        println("=> range map by key", m[k])
        //=> range map by key zero
        //=> range map by key one
    }

    for k, v := range m {
        println("=> range map by key-value", k, v)
        //=> range map by key-value 0 zero
        //=> range map by key-value 1 one
    }

    for _, v := range m {
        println("=> range map by value", v)
        //=> range map by value zero
        //=> range map by value one
    }

    switch m["0"] {
    case "two", "three":
        println("=> switch:", "two or three")
        // We don't need 'break' to prevent fallthrough
        fallthrough //!!!
    case "four":
        println("=> switch:", "four")
        if m["0"] == "one" {
            break
        }
        fallthrough //!!!
    case "zero", "five":
        println("=> switch:", "zero or five")
        fallthrough
    default:
        println("=> switch: some other number")
    }
    //=> switch: zero or five
    //=> switch: some other number

    for _, v := range m {
        switch v {
        case "zero":
            break
        default:
            println("=> switch: some other number")
        }
        break
    }
    //no output
}
