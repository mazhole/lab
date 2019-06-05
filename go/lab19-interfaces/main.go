package main

import "fmt"

type Flier interface {
    Fly()
}

type Bird struct {
    Name string
}

func (b Bird) Fly() {
    fmt.Println("=> Bird.Fly:", b.Name)
}

func DoFly(f Flier) {
    f.Fly()
}

func GoFly(f Flier) {
    f.Fly()
    //=> Bird.Fly: Bee

    // If you are sure that f is exactly Bird otherwise PANIC
    //b := f.(Bird)

    // Compilation error: f.Name undefined (type Flier has no field or method Name)
    //f.Name = "Hi!"

    //Type assertion
    if b, ok := f.(Bird); ok {
        b.Fly()
        //=> Bird.Fly: Bee
    }
}

func main() {
    b := Bird {"Bee"}
    DoFly(b)
    //=> Bird.Fly: Bee

    GoFly(b)
}
