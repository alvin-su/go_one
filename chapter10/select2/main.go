package main

import "fmt"

func main() {
	ch1 := make(chan string)
	go makeFoods(ch1)
	for val := range ch1 {
		fmt.Println(val + "菜好了")
	}
	fmt.Println("您的菜上齐了!")
}

func makeFoods(c chan string) {
	foods := []string{"红烧肉", "油焖大虾", "水煮肉片"}
	for _, val := range foods {
		c <- val
	}
	close(c)
}
