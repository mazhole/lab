package main

import (
    "fmt"
    "sync"
)

const defaultTaskChannelCapacity = 1

type Task interface {
    Execute()
}

type Pool struct {
    mu    sync.Mutex
    size  int
    tasks chan Task
    kill  chan struct{}
    wg    sync.WaitGroup
}

func NewPool(size int) *Pool {
    pool := &Pool {
        tasks: make(chan Task, defaultTaskChannelCapacity),
        kill: make(chan struct{}),
    }
    pool.Resize(size)
    return pool
}

func (p *Pool) worker() {
    fmt.Println("=> Pool.worker: worker is started")
    defer p.wg.Done()
    for {
        select {
        case task, ok := <-p.tasks:
            if !ok {
                fmt.Println("=> Pool.worker: Can't get the task from tasks channel")
                return
            }
            task.Execute()
        case <-p.kill:
            fmt.Println("=> Pool.worker: worker is killed")
            return
        }
    }
}

func (p *Pool) Resize(n int) {
    p.mu.Lock()
    defer p.mu.Unlock()
    for p.size < n {
        p.size++
        p.wg.Add(1)
        go p.worker()
    }
    for p.size > n {
        p.size--
        p.kill <- struct{}{}
    }
    fmt.Println("=> Pool.size:", p.size)
}

func (p *Pool) Close() {
    p.Resize(0)
    close(p.tasks)
}

func (p *Pool) Wait() {
    p.wg.Wait()
}

func (p *Pool) Exec(task Task) {
    p.tasks <- task
}

type ExampleTask string

func (et ExampleTask) Execute() {
    fmt.Println("=> Task executing:", string(et))
}

func main() {
    pool := NewPool(5)

    pool.Exec(ExampleTask("1"))
    pool.Exec(ExampleTask("2"))

    pool.Resize(3)
    pool.Resize(6)

    for i := 0; i < 20; i++ {
        pool.Exec(ExampleTask(fmt.Sprintf("additional task %d", i + 3)))
    }

    pool.Close()

    pool.Wait()

    fmt.Println("=> Done!")
}

//Output:
//=> Pool.size: 5
//=> Pool.worker: worker is started
//=> Task executing: 1
//=> Task executing: 2
//=> Pool.worker: worker is started
//=> Pool.worker: worker is started
//=> Pool.worker: worker is started
//=> Pool.worker: worker is killed
//=> Pool.worker: worker is killed
//=> Pool.size: 3
//=> Pool.worker: worker is started
//=> Pool.size: 6
//=> Task executing: additional task 5
//=> Task executing: additional task 6
//=> Task executing: additional task 7
//=> Task executing: additional task 8
//=> Pool.worker: worker is started
//=> Task executing: additional task 3
//=> Task executing: additional task 10
//=> Task executing: additional task 11
//=> Pool.worker: worker is started
//=> Pool.worker: worker is started
//=> Task executing: additional task 9
//=> Task executing: additional task 12
//=> Task executing: additional task 14
//=> Task executing: additional task 17
//=> Task executing: additional task 13
//=> Task executing: additional task 18
//=> Task executing: additional task 16
//=> Task executing: additional task 15
//=> Task executing: additional task 4
//=> Pool.worker: worker is killed
//=> Task executing: additional task 19
//=> Pool.worker: worker is killed
//=> Task executing: additional task 20
//=> Pool.worker: worker is killed
//=> Pool.worker: worker is killed
//=> Task executing: additional task 21
//=> Pool.worker: worker is killed
//=> Task executing: additional task 22
//=> Pool.worker: worker is killed
//=> Pool.size: 0
//=> Done!