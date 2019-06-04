package main

import "fmt"

type Person struct {
    Age int
    Name string
}

type Employee struct {
    Age int
}

type SecretAgent1 struct {
    Person   // Is embedded into SecretAgent
    Employee // Is embedded into SecretAgent
    LicenseToKill bool
}

type SecretAgent2 struct {
    Name string
    Person   // Is embedded into SecretAgent
    Employee // Is embedded into SecretAgent
    LicenseToKill bool
}

func (p Person) GetName() string {
    return p.Name
}

func (s SecretAgent2) GetName() string {
    return s.Name
}

// Compilation error!!!
//func (s SecretAgent2) GetName() int {
//    return s.Name
//}

func main() {
    sa1 := SecretAgent1 {
        Person: Person {
            Age: 27,
            Name: "Jack",
        },
        LicenseToKill: true,
    }

    sa2 := SecretAgent2 {
        Name: "Jeremy",
        Person: Person {
            Age: 33,
            Name: "John",
        },
        Employee: Employee {
            Age: 35,
        },
    }

    fmt.Printf("=> sa1: %T %+v\n", sa1, sa1)
    //=> sa1: main.SecretAgent1 {Person:{Age:27 Name:Jack} Employee:{Age:0} LicenseToKill:true}

    fmt.Printf("=> sa2: %T %+v\n", sa2, sa2)
    //=> sa2: main.SecretAgent2 {Name:Jeremy Person:{Age:33 Name:John} Employee:{Age:35} LicenseToKill:false}

    fmt.Println("=> sa1.Person.Name:", sa1.Person.GetName())
    //=> sa1.Person.Name: Jack

    fmt.Println("=> sa1.Name:", sa1.GetName())
    //=> sa1.Name: Jack

    fmt.Println("=> sa2.Name:", sa2.GetName())
    //=> sa2.Name: Jeremy
}
