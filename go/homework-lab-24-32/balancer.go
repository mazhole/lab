package main

import (
//    "fmt"
//    "reflect"
    "sync"
//    "testing"
)

var nextNodeIndex int

var mutex sync.RWMutex

type Node struct {
    Statistic int
}

type RoundRobinBalancer struct {
    Nodes []Node
}

// Initialization
func (b *RoundRobinBalancer) Init(n int) {
    nextNodeIndex = 0
    b.Nodes = make([]Node, n)
}

// Returns quantity of requests by every node
func (b *RoundRobinBalancer) GetStatistic() (statistic []int) {
    statistic = make([]int, len(b.Nodes))
    for i, n := range b.Nodes {
        mutex.RLock()
        defer mutex.RUnlock()
        statistic[i] = n.Statistic
    }
    return statistic
}

// Returns the server to process the request
func (b *RoundRobinBalancer) GetNode() (node Node) {
    mutex.Lock()
    defer mutex.Unlock()
    b.Nodes[nextNodeIndex].Statistic++
    node = b.Nodes[nextNodeIndex]
    if (nextNodeIndex == len(b.Nodes) - 1) {
        nextNodeIndex = 0
    } else {
        nextNodeIndex++
    }
    return node
}