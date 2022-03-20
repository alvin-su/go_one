package main

import "fmt"

func main() {
	var list [3]string
	fmt.Println(list)
	var list2 = [3]string{"string1", "string2", "string3"}
	fmt.Println(list2)
	list3 := [...]string{"test1", "test2", "test2"}
	fmt.Println(list3)
	for idx, item := range list3 {
		//拼接字符串，将结果放到 desc 变量中
		desc := fmt.Sprintf("%d-%s", idx, item)
		fmt.Println(desc)
	}
}
