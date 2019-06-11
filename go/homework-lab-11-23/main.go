package main

import "fmt"

// Decorator

type actionInfo struct {
    data string
}

type worker interface {
    Do(*actionInfo) (*actionInfo)
}

type action func(*actionInfo) (*actionInfo)

func (a action) Do(ai *actionInfo) (*actionInfo) {
    return a(ai)
}

type decorator func(worker) worker

//func (d decorator) decorate(w worker) worker
//{
//    return d(w)
//}

func decorate(w worker, ds ...decorator) worker {
    fmt.Println("=> decorate: decorators length:", len(ds))

    decoratedWorker := w
    for _, d := range ds {
        decoratedWorker = d(decoratedWorker)
    }
    return decoratedWorker
}


type fileManager struct {
}

func (fileManager) Do(ai *actionInfo) (*actionInfo) {
    fmt.Println("=> fileManager.Do func")
    fmt.Println("=> I am FileManager:", ai.data)
    return ai
}


func read(fileName string) decorator {
    fmt.Println("=> read func")
    return func(w worker) worker {
        fmt.Println("=> read worker func")
        return action(func(ai *actionInfo) (*actionInfo) {
            fmt.Println("=> read action func:", fileName, "Input:", ai.data)
            ai.data = "read: done. Let's write!"
            return w.Do(ai)
        })
    }
}

func write(fileName, flag string) decorator {
    fmt.Println("=> write func")
    return func(w worker) worker {
        fmt.Println("=> write worker func")
        return action(func(ai *actionInfo) (*actionInfo) {
            fmt.Println("=> write action func:", fileName, flag, "Input:", ai.data)
            ai.data = "write: done. Let's print!"
            return w.Do(ai)
        })
    }
}

func print(fileName, printerName string) decorator {
    fmt.Println("=> print func")
    return func(w worker) worker {
        fmt.Println("=> print worker func")
        return action(func(ai *actionInfo) (*actionInfo) {
            fmt.Println("=> print action func::", fileName, printerName, "Input:", ai.data)
            ai.data = "print: done!"
            return w.Do(ai)
        })
    }
}

func main() {
    ai := actionInfo { "empty" }

    fm := fileManager{}

    decorators := []decorator {
        print("test2.txt", "printer1"),
        write("test2.txt", "flag"),
        read("test1.txt"),
        
        
    }

    rr := decorate(&fm, decorators...)

    rr.Do(&ai)
}

//=> print func
//=> write func
//=> read func
//=> decorate: decorators length: 3
//=> print worker func
//=> write worker func
//=> read worker func
//=> read action func: test1.txt Input: empty
//=> write action func: test2.txt flag Input: read: done. Let's write!
//=> print action func:: test2.txt printer1 Input: write: done. Let's print!
//=> fileManager.Do func
//=> I am FileManager: print: done!