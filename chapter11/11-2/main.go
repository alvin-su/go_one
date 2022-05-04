package main

import (
	"fmt"
	"sync"
	"time"
)

//var m sync.Mutex //互斥锁
var m sync.RWMutex //读写锁

func main() {
	s := []string{"老张", "老王", "老李"}
	for i := 0; i < len(s); i++ {
		//go GetSeatNoLock(s[i])
		go CheckSeat(s[i])
		go GetSeat(s[i])
	}
	//time.Sleep(5 * time.Second)
	time.Sleep(8 * time.Second)
}

func GetSeatNoLock(name string) {
	fmt.Println(name + "已经抢到位置了。")
	time.Sleep(1 * time.Second)
	fmt.Println(name + "已经离开。")
}

func GetSeat(name string) {
	m.Lock()
	defer m.Unlock()
	fmt.Println(name + "已经抢到位置。")
	time.Sleep(1 * time.Second)
	fmt.Println(name + "已经离开。")
}

// CheckSeat 读锁
func CheckSeat(name string) {
	m.RLock()
	defer m.RUnlock()
	fmt.Println(name + "查看位置。")
	time.Sleep(1 * time.Second)
}
