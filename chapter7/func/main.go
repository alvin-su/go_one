package main

import (
	"fmt"
	"strconv"
	"sync"
)

func main() {
	todayFoods := [3]string{"卤肉饭", "西红柿鸡蛋饭", "肥牛饭"}
	ShowFoods(todayFoods)
	fmt.Println(todayFoods)

	todayFoods2 := &[3]string{"卤肉饭", "西红柿鸡蛋饭", "肥牛饭"}
	ShowFoods2(todayFoods2)
	fmt.Println(todayFoods2)

	todayFoods3 := []string{"卤肉饭", "西红柿鸡蛋饭", "肥牛饭"}
	ShowFoods3(todayFoods3)
	fmt.Println(todayFoods3)

	result := Desc("粤菜", "盐焗鸡")
	fmt.Println(result)

	m := Menu("红烧肉", 88)
	fmt.Println(m)

	var hello Hi        //定义函数变量 hello
	hello = Hello       //将函数 Hello 赋值给函数变量 hello
	words := hello("3") //运行Hello 函数
	fmt.Printf("%s\n", words)

	hello = Hello4DongBei
	words = hello("5")
	fmt.Printf("%s\n", words)

	SayHello("3", hello) //函数变量 hello 作为参数传递给 函数

	//匿名函数
	hello2 := func(num string) string {
		return num + "位兄弟，欢迎光临"
	}

	hello3 := func(num string) string {
		return num + "位客人，欢迎光临"
	}

	SayHello("10", hello2)
	SayHello("20", hello3)

	// 闭包
	f := PayOrder(func() float64 {
		return 0.8
	})

	result2 := f("红烧肉", 88)
	fmt.Println(result2)

	m["abc"] = 123
	m["m"] = 3
	m["g"] = 80
	lookup("g")

	start()
	testFood()
	end()
}

// ShowFoods
//  ShowFoods
//  @Description: 函数传递数组
//  @param foods 数组
//
func ShowFoods(foods [3]string) {
	foods[0] = "蛋炒饭"
	fmt.Println(foods)
}

// ShowFoods2 传递数组指针
func ShowFoods2(foods *[3]string) {
	foods[0] = "蛋炒饭"
	fmt.Println(foods)
}

// ShowFoods3 传递切片
func ShowFoods3(foods []string) {
	foods[0] = "蛋炒饭"
	fmt.Println(foods)
}

// Desc 函数返回值
func Desc(foodType, foodName string) string {
	return "这是" + foodType + "的" + foodName
}

// Menu 返回字典
func Menu(name string, price int) map[string]int {
	m := make(map[string]int)
	m[name] = price
	return m
}

// Hi 函数类型的声明
type Hi func(num string) string

func Hello(num string) string {
	return num + "位客人，欢迎光临"
}

func Hello4DongBei(num string) string {
	return num + "位兄弟，欢迎光临"
}

func SayHello(num string, hi Hi) {
	result := hi(num)
	fmt.Println(result)
}

type Discount func() float64
type CheckSum func(name string, price float64) float64

func PayOrder(discount Discount) CheckSum {
	var total float64
	return func(name string, price float64) float64 {
		fmt.Println("菜品名称:" + name + "单价:" + strconv.FormatFloat(price, 'f', -1, 64))
		total = total + price
		if discount == nil {
			return total
		} else {
			return total * discount()
		}
	}
}

func Total(prices ...int) int {
	result := 0
	for _, val := range prices {
		result += val
	}
	return result
}

// 保证使用映射时的并发安全的互斥锁
var mu sync.Mutex
var m = make(map[string]int)

func lookup(key string) int {
	mu.Lock()
	defer mu.Unlock()
	return m[key]
}

func start() {
	fmt.Println("程序开始执行...")
}

func testFood() {
	foods := []string{"123", "456", "789"}
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("defer finished")
		}
	}()
	fmt.Println(foods[3])
}

func end() {
	fmt.Println("程序执行结束...")
}
