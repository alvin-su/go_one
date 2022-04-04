package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	message := "hello Go world"
	fmt.Println(message)
	message = "Hello Go world,Wonderfully"
	fmt.Println(message)

	name := "peter"
	fmt.Println(strings.ToUpper(name))
	//字符串拼接的多种方式

	phone := "1300000000"
	user := name + " " + phone
	fmt.Println(user)
	user1 := strings.Join([]string{"hellow", "world"}, "-")
	fmt.Println(user1)
	user2 := fmt.Sprintf("%s:%s", name, phone)
	fmt.Println(user2)
	user3 := bytes.Buffer{}
	user3.WriteString("hello")
	user3.WriteString(" world")
	fmt.Println(user3.String())

	//删除空格
	foods3 := "   米饭，馒头  "
	foods3 = strings.Trim(foods3, " ")
	fmt.Println(foods3)

	age := 25
	fmt.Println("我的年龄是" + strconv.Itoa(age))

	age2 := 36
	fmt.Println("我的年龄是" + strconv.Itoa(age2))
}
