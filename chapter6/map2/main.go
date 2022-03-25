package main

import (
	"fmt"
	"sort"
)

func main() {
	//用内置的make函数创建 map
	ages := make(map[string]int)
	ages["张三"] = 18
	ages["李四"] = 20
	fmt.Println(ages)

	//使用字面值语法创建map
	ages2 := map[string]int{"张三": 30, "李四": 40}
	fmt.Println(ages2)

	//创建空的map表达式语法
	ages3 := map[string]int{}
	ages3["张三"] = 50
	ages3["李四"] = 60
	fmt.Println(ages3)
	//Map中的元素通过key对应的下标语法访问
	fmt.Println(ages3["张三"])

	//使用内置的delete函数删除元素
	delete(ages3, "张三")
	fmt.Println(ages3["张三"]) //被删除了，返回int类型的 零值

	//要想遍历map中全部的key/value对的话，可以使用range风格的for循环实现，和之前的slice遍历语法类似。
	//下面的迭代语句将在每次迭代时设置name和age变量
	for name, age := range ages {
		fmt.Printf("%s\t%d\n", name, age)
	}

	//要按顺序遍历key/value对，我们必须显式地对key进行排序，可以使用sort包的Strings函数对字符串slice进行排序

	var names []string
	for name := range ages2 {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, ages2[name])
	}

	//判断map中是否存在某个k,v
	age, ok := ages["李四"]
	if !ok {
		fmt.Println("map：ages中不存在王五")
	} else {
		fmt.Println(age)
	}

	if _, ok := ages["王五"]; !ok {
		fmt.Println("map：ages中不存在王五")
	}

	//比较两个map 是否相等
	var isEqual = equal(map[string]int{"A": 2}, map[string]int{"A": 2})
	fmt.Println(isEqual) //true

	var isEqual2 = equal(ages, ages2)
	fmt.Println(isEqual2) //false
}

//比较两个map 是否相等
func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, xv := range x {
		if yv, ok := y[k]; !ok || xv != yv {
			return false
		}
	}
	return true
}
