package main

import "fmt"

type account struct {
    balance float64
}

func (a *account) Credit(amount float64) {
    a.balance += amount
    fmt.Printf("=> Credit. Amount: %f Balance: %f\n", amount, a.balance)
}

func (a *account) Debit(amount float64) {
    a.balance -= amount
    fmt.Printf("=> Debit. Amount: %f Balance: %f\n", amount, a.balance)
}

func main() {
    a := account{}

    for i := 0; i < 20; i++ {
        go func() {
            a.Debit(float64(i))
        }()
        go func() {
            a.Credit(float64(i))
        }()
    }

    fmt.Scanln()

    fmt.Printf("=> Balance: %f\n", a.balance)
}

//Output:
//=> Debit. Amount: 1.000000 Balance: -1.000000
//=> Credit. Amount: 14.000000 Balance: 17.000000
//=> Credit. Amount: 20.000000 Balance: 77.000000
//=> Credit. Amount: 20.000000 Balance: 57.000000
//=> Debit. Amount: 20.000000 Balance: 37.000000
//=> Credit. Amount: 20.000000 Balance: 77.000000
//=> Credit. Amount: 20.000000 Balance: 57.000000
//=> Debit. Amount: 20.000000 Balance: 37.000000
//=> Credit. Amount: 20.000000 Balance: 57.000000
//=> Credit. Amount: 20.000000 Balance: 57.000000
//=> Debit. Amount: 10.000000 Balance: -7.000000
//=> Credit. Amount: 10.000000 Balance: 3.000000
//=> Debit. Amount: 20.000000 Balance: 37.000000
//=> Debit. Amount: 14.000000 Balance: -11.000000
//=> Credit. Amount: 20.000000 Balance: 77.000000
//=> Credit. Amount: 14.000000 Balance: 3.000000
//=> Debit. Amount: 14.000000 Balance: -11.000000
//=> Debit. Amount: 20.000000 Balance: 57.000000
//=> Credit. Amount: 14.000000 Balance: 3.000000
//=> Debit. Amount: 14.000000 Balance: -11.000000
//=> Credit. Amount: 20.000000 Balance: 77.000000
//=> Credit. Amount: 20.000000 Balance: 57.000000
//=> Credit. Amount: 14.000000 Balance: 3.000000
//=> Debit. Amount: 20.000000 Balance: 37.000000
//=> Debit. Amount: 14.000000 Balance: -11.000000
//=> Credit. Amount: 14.000000 Balance: 3.000000
//=> Credit. Amount: 20.000000 Balance: 57.000000
//=> Debit. Amount: 20.000000 Balance: 57.000000
//=> Debit. Amount: 14.000000 Balance: -11.000000
//=> Debit. Amount: 20.000000 Balance: 37.000000
//=> Credit. Amount: 14.000000 Balance: 3.000000
//=> Credit. Amount: 20.000000 Balance: 37.000000
//=> Credit. Amount: 20.000000 Balance: 97.000000
//=> Debit. Amount: 20.000000 Balance: 77.000000
//=> Debit. Amount: 20.000000 Balance: 57.000000
//=> Debit. Amount: 6.000000 Balance: -7.000000
//=> Credit. Amount: 10.000000 Balance: 3.000000
//=> Debit. Amount: 20.000000 Balance: 37.000000
//=> Debit. Amount: 20.000000 Balance: 17.000000
//=> Debit. Amount: 20.000000 Balance: 57.000000
//
//=> Balance: 17.000000