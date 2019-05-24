package main

import "fmt"

func main() {
    var symbol rune = 'a'
    var autoSymbol = 'a'
    unicodeSymbol := '~'
    unicodeSymbolByNumber := '\u2318'

    fmt.Println("=> fmt.Println:", symbol, autoSymbol, unicodeSymbol, unicodeSymbolByNumber)
    //=> fmt.Println: 97 97 126 8984

    s1 := "Тест!"
    fmt.Println("=> s1 (ru):", s1, len(s1))
    //=> s1 (ru): Тест! 9
    for i, runeValue := range s1 {
        fmt.Printf("fmt.Printf: %#U at position %d\n", runeValue, i)
    }
    //fmt.Printf: U+0422 'Т' at position 0
    //fmt.Printf: U+0435 'е' at position 2
    //fmt.Printf: U+0441 'с' at position 4
    //fmt.Printf: U+0442 'т' at position 6
    //fmt.Printf: U+0021 '!' at position 8

    s2 := "法庭文件发"
    fmt.Println("=> s2 (ru):", s2, len(s2))
    //=> s2 (ru): 法庭文件发 15
    for i, runeValue := range s2 {
        fmt.Printf("fmt.Printf: %#U at position %d\n", runeValue, i)
    }
    //fmt.Printf: U+6CD5 '法' at position 0
    //fmt.Printf: U+5EAD '庭' at position 3
    //fmt.Printf: U+6587 '文' at position 6
    //fmt.Printf: U+4EF6 '件' at position 9
    //fmt.Printf: U+53D1 '发' at position 12

    fmt.Println("=> s2[14]:", s2[14])
    //=> s2[14]: 145

    bin := []byte(s2)
    fmt.Println("=> binary cn:", bin, len(bin))
    //=> binary cn: [230 179 149 229 186 173 230 150 135 228 187 182 229 143 145] 15
    for i, v := range bin {
        fmt.Printf("=> raw binary i: %v, oct: %v, hex: %x\n", i, v, v)
    }
    //=> raw binary i: 0, oct: 230, hex: e6
    //=> raw binary i: 1, oct: 179, hex: b3
    //=> raw binary i: 2, oct: 149, hex: 95
    //=> raw binary i: 3, oct: 229, hex: e5
    //=> raw binary i: 4, oct: 186, hex: ba
    //=> raw binary i: 5, oct: 173, hex: ad
    //=> raw binary i: 6, oct: 230, hex: e6
    //=> raw binary i: 7, oct: 150, hex: 96
    //=> raw binary i: 8, oct: 135, hex: 87
    //=> raw binary i: 9, oct: 228, hex: e4
    //=> raw binary i: 10, oct: 187, hex: bb
    //=> raw binary i: 11, oct: 182, hex: b6
    //=> raw binary i: 12, oct: 229, hex: e5
    //=> raw binary i: 13, oct: 143, hex: 8f
    //=> raw binary i: 14, oct: 145, hex: 91

    //useful package:
    //https://golang.org/src/unicode/utf8/utf8.go
}
