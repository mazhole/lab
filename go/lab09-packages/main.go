package main

import w "github.com/mazhole/study/go/lab09-packages/world"

func main() {
    //world.PrintStartRoom()

    //alias
    w.PrintStartRoom()
    //Start level is lobby

    //println("=> starting level:", w.startingLevel)
    //./main.go:13:35: cannot refer to unexported name world.startingLevel
    //./main.go:13:35: undefined: world.startingLevel

    println("=> starting level:", w.GetStartingLevel())
    //=> starting level: 1
}
