package main

import "fmt"

func main() {
	foods := []string{"红烧肉", "清蒸鱼", "熘大虾", "蒸螃蟹", "鲍鱼粥"}
	for _, item := range foods {
		if item == "鲍鱼粥" {
			fmt.Println("鲍鱼粥,今日免费送")
		} else {
			fmt.Println(item)
		}
	}

	prices := [...]int16{32, 68, 96, 153, 198, 77, 100}
	switch prices[1] {
	case 32, 68:
		fmt.Println("0 or 1")
	case 96, 153:
		fmt.Println("2 or 3")
	case 198, 77, 100:
		fmt.Println("4 or 5 or 6")

	}
}
