package main

import (
    "fmt"
    "sort"
)

type MyInt int

func (m MyInt) showYourSelf() {
    fmt.Printf("%T %v\n", m, m)
}

func (m *MyInt) add(i MyInt) {
    *m = *m + MyInt(i)
}

type MyStruct struct {
    No   int
    Name string
}

type MySliceStruct []MyStruct

func (m MySliceStruct) Less(i int) bool {
    return len(m) < i
}

func (m MySliceStruct) Len() int {
    return len(m)
}

func (m MySliceStruct) Swap(i, j int) {
    t := m[i]
    m[i] = m[j]
    m[j] = t
}

func main() {
    myInt := MyInt(0)
    myInt.add(3)
    myInt.showYourSelf()

    mss := MySliceStruct {{
            No: 1,
            Name: "1",
        },{
            No: 9,
            Name: "9",
        },{
            No: 2,
            Name: "2",
        },{
            No: 8,
            Name: "8",
        },{
            No: 3,
            Name: "3",
        },{
            No: 7,
            Name: "7",
        },{
            No: 4,
            Name: "4",
        },{
            No: 6,
            Name: "6",
        },{
            No: 5,
            Name: "5",
        },
    }
    fmt.Println("=> ", mss)
    //=>  [{1 1} {9 9} {2 2} {8 8} {3 3} {7 7} {4 4} {6 6} {5 5}]

    sorter(mss)
    fmt.Println("=> ", mss)
    //=>  [{1 1} {2 2} {3 3} {4 4} {5 5} {6 6} {7 7} {8 8} {9 9}]

    mss[8].Name = "Test9"
    
    mss.Swap(2, 4)
    fmt.Println("=> ", mss, mss.Less(5), mss.Len())
    //=>  [{1 1} {2 2} {5 5} {4 4} {3 3} {6 6} {7 7} {8 8} {9 Test9}] false 9
}

func sorter(sl []MyStruct) []MyStruct {
    sort.Slice(sl[:], func(i, j int) bool { return sl[i].No < sl[j].No })
    return sl
}
