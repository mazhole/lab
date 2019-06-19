package main

import (
    "fmt"
    "runtime"
    "sync"
    "sync/atomic"
)

const stepsCount = 1000

type account struct {
    amount int64
    wrongAmount int64
}

func (a *account) Balance() int64 {
    return atomic.LoadInt64(&a.amount)
}

func (a *account) Credit(amount int64) {
    atomic.AddInt64(&a.amount, amount)
}

func (a *account) WrongCredit(amount int64) {
    a.wrongAmount += amount
}

func (a *account) Debit(amount int64) {
    atomic.AddInt64(&a.amount, -1 * amount)
}

func main() {
    a := account{ amount: 0 }

    var wg sync.WaitGroup
    wg.Add(stepsCount)

    for i := 0; i < stepsCount; i++ {
        go func(wg *sync.WaitGroup) {
            for i := 0; i < stepsCount; i++ {
                a.Credit(1)
                a.WrongCredit(1)
                runtime.Gosched()
            }
            wg.Done()
        }(&wg)
    }

    wg.Wait()

    fmt.Printf("=> Right result: %d\n", a.Balance())
    //=> Right result: 1000000

    fmt.Printf("=> Wrong resilt: %d\n", a.wrongAmount)
    //=> Wrong result: 988472
}