package main

import (
	"fmt"
	"sync"
	"time"
)

var lock sync.RWMutex
var buySth string
var box sync.RWMutex
var sendCond = sync.NewCond(&box)
var recvCond = sync.NewCond(box.RLocker())

func main() {
	go send()
	go receive()
	time.Sleep(5 * time.Second)
}

func send() {

	box.Lock()
	for buySth == "已投递" {
		sendCond.Wait()
	}
	buySth = "已投递"
	box.Unlock()
	recvCond.Signal()
}

func receive() {
	lock.RLock()
	for buySth == "" {
		recvCond.Wait()
	}
	buySth = ""
	fmt.Println("货物已取走")
	lock.RUnlock()
	sendCond.Signal()
}
