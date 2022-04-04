package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println(i) //不会打印任何东西
		}()
	}
}
