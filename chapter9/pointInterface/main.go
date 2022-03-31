package main

import "fmt"

func main() {
	li := Chef{
		Name: "李师傅",
		Age:  60,
	}
	var i IChef = &li
	fmt.Println(i.Cook())
	fmt.Println(i.FavCook("盐焗鸡"))
}

type IChef interface {
	Cook() string
	FavCook(foodName string) string
}

type Chef struct {
	Name string
	Age  int
}

func (c *Chef) Cook() string {
	return c.Name + "饭菜做好了"
}

func (c *Chef) FavCook(foodName string) string {
	return c.Name + "的拿手好菜" + foodName + "做好了"
}
