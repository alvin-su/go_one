package main

import "fmt"

func main() {
	li := Chef{
		Name: "李师傅",
		Age:  50,
	}
	var inter IChef = li
	fmt.Println(inter.Cook())
	fmt.Println(inter.FavCook("辣椒炒肉"))

	//空接口调用，空接口可以接收任意类型的变量
	var i interface{}
	Say(i)
	Say(77)
	Say("Go语言学习")
}

type IChef interface {
	Cook() string
	FavCook(foodName string) string
}

type Chef struct {
	Name string
	Age  int
}

func (c Chef) Cook() string {
	result := c.Name + "饭菜做好了"
	return result
}

func (c Chef) FavCook(foodName string) string {
	result := c.Name + "的拿手好菜:" + foodName + "做好了"
	return result
}

// Say 空接口
func Say(i interface{}) {
	fmt.Printf("(%v,%T)\n", i, i)
}
