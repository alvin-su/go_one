package main

import "fmt"

var balance int

func main() {
	for i := 0; i < 10; i++ {
		//Sum(i)
		//并发执行
		go Sum(i)
	}
	fmt.Println(GetTotal())
}

func Sum(amount int) {
	balance = balance + amount
}

func GetTotal() int {
	return balance
}
