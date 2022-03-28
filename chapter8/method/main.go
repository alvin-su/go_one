package main

import "fmt"

func main() {
	li := Chef3{
		Name: "李师傅",
		Age:  30,
	}
	result := li.Cook("红烧肉")
	fmt.Println(result)
	result = li.FavCook("萝卜干腊肉")
	fmt.Println(result)

	li2 := &Chef3{
		Name: "张师傅",
		Age:  40,
	}
	result2 := li2.CookPoint("西红柿炒蛋")
	fmt.Println(result2)

	li3 := Chef3{
		Name: "王师傅",
		Age:  50,
	}
	liPoint := &li3
	result3 := liPoint.CookPoint("烤牛油")
	fmt.Println(result3)
}

type Chef3 struct {
	Name string
	Age  int
}

func (c Chef3) Cook(name string) string {
	return c.Name + "：做好了 " + name
}

func (c Chef3) FavCook(name string) string {
	return c.Name + "：这是我的拿手好菜" + name
}

// CookPoint 指针方法
func (c *Chef3) CookPoint(name string) string {
	return c.Name + "：做好了 " + name
}
