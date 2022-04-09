package main

import "fmt"

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	go first(ch1)
	go second(ch1, ch2)
	Cook(ch2)
	fmt.Println("洗碗...")
}

func first(c chan<- string) {
	c <- "买菜"
	close(c)
}

func second(c1 <-chan string, c2 chan<- string) {
	r := <-c1
	c2 <- r + "买肉"
	close(c2)
}

func Cook(c <-chan string) {
	r := <-c
	fmt.Println(r + "已经准备好了，吃顿好的！")
}
