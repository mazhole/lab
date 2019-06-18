package main

import (
    "fmt"
    //"runtime"
    "sync"
)

const stepsCount = 5

type account struct {
    sync.RWMutex
    balance int
}

func (a *account) Balance() int {
    a.RLock()
    defer a.RUnlock()
    return a.balance
}

func (a *account) Credit(amount int) {
    a.Lock()
    defer a.Unlock()
    a.balance += amount
    fmt.Printf("=> Credit. Amount: %d Balance: %d\n", amount, a.balance)
}

func (a *account) Debit(amount int) {
    for {
        if a.balance >= amount {
            a.Lock()
            if a.balance >= amount {
                a.balance -= amount
                fmt.Printf("=> Debit. Amount: %d Balance: %d\n", amount, a.balance)
                a.Unlock()
                return
            } else {
                a.Unlock()
            }
        }
    }
}

func main() {
    var wg sync.WaitGroup
    wg.Add(2 * stepsCount)
    a := account{}

    for i := 0; i < stepsCount; i++ {
        go func(amount int, wg *sync.WaitGroup) {
            a.Debit(amount)
            wg.Done()
        }(i, &wg)
        go func(amount int, wg *sync.WaitGroup) {
            a.Credit(amount)
            wg.Done()
        }(i, &wg)
    }

    wg.Wait()

    fmt.Printf("=> Balance: %d\n", a.balance)
}

//Output:
//=> Credit. Amount: 4 Balance: 4
//=> Debit. Amount: 3 Balance: 1
//=> Credit. Amount: 1 Balance: 2
//=> Credit. Amount: 3 Balance: 5
//=> Debit. Amount: 0 Balance: 5
//=> Credit. Amount: 0 Balance: 5
//=> Debit. Amount: 1 Balance: 4
//=> Debit. Amount: 4 Balance: 0
//=> Credit. Amount: 2 Balance: 2
//=> Debit. Amount: 2 Balance: 0
//=> Balance: 0