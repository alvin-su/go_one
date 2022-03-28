package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Chef struct {
	Name string
	Age  int
	Honor
	Trainee []Chef
}

type Honor struct {
	Title   string
	GetTime time.Time
}

func main() {
	wang := Chef{"王师傅", 30, Honor{}, nil}
	fmt.Println(wang.Name)
	fmt.Println((&wang).Name)

	li := Chef{
		Name:    "李师傅",
		Age:     35,
		Honor:   Honor{},
		Trainee: nil,
	}
	fmt.Println(li.Name)

	li2 := &Chef{
		Name:    "张师傅",
		Age:     40,
		Honor:   Honor{},
		Trainee: nil,
	}
	fmt.Println(li2.Name)
	li2.Title = "中华厨师奖"
	fmt.Println(li2.Title)
	chef := struct {
		Name string
		Age  int
	}{
		Name: "老王",
		Age:  58,
	}

	fmt.Println(chef.Name)

	list := []string{"张三", "李四", "王五"}
	//showName(list)
	fmt.Println(list)
	showName2(list)
	fmt.Println(list)

	c := Chef2{
		Name: "小李",
		Age:  18,
	}
	marshal, err := json.Marshal(&c)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(marshal))

	var cc Chef2
	s := `{"name":"小张","Age":20}`
	err = json.Unmarshal([]byte(s), &cc)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(cc.Name)
	fmt.Println(cc.Age)
}

func showName(names []string) {
	for _, name := range names {
		fmt.Println("这是：" + name)
	}
}

func showName2(names []string) {
	for idx, name := range names {
		if name == "王五" {
			names[idx] = "隔壁老王"
		}
		fmt.Println("这是：" + name)
	}
}

type Chef2 struct {
	Name string
	Age  int
}
