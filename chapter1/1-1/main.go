package main

import "fmt"

func main() {
	//fmt.Println("Hello Go world")

	var name string = "Hello C# world"
	fmt.Println(name)
	fmt.Println(&name)
	p := &name
	fmt.Println(p)
	fmt.Println(*p)
	*p = "Hello Go world"
	fmt.Println(*p)
	fmt.Println(name)

	name2 := name + *new(string)
	fmt.Println(name2)
}
