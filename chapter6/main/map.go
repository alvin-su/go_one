package main

import "fmt"

func main() {
	//一、创建字典
	//1.1 使用字面量创建字典
	var m map[string]int = map[string]int{"红烧肉": 88, "清蒸鱼": 98, "熘大虾": 128}
	fmt.Println(m)

	//1.2 使用内置make函数创建字典
	foodsMap := make(map[string]int)
	foodsMap["红烧肉"] = 88
	foodsMap["清蒸鱼"] = 98
	foodsMap["熘大虾"] = 128
	fmt.Println(foodsMap)

	//二、使用字典
	addressList := make(map[string]int)
	addressList["张三"] = 123
	addressList["李四"] = 456
	fmt.Println(addressList)
	addressList["张三"] = 123456
	fmt.Println(addressList)
	fmt.Println(addressList["李四"])
	addressList["王五"] = 789
	i, ok := addressList["王五"]
	if !ok {
		fmt.Println("addressList没有找到对应的键")
	} else {
		fmt.Println(i)
	}
	for idex, item := range addressList {

		fmt.Println(idex)
		fmt.Println(item)
	}

	fmt.Println(addressList)
	delete(addressList, "王五")
	fmt.Println(addressList)
}
