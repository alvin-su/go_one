package main

import "fmt"

//遍历切片元素
func sliceForach(slice []string) {
	for idx, item := range slice {
		fmt.Println(idx)
		fmt.Println(item)
		fmt.Println("------")
	}
}
