package main

import (
	"fmt"
	"reflect"
)

type T struct{}

func main() {
	name := "GoodDinner"
	t := &T{}
	reflect.ValueOf(t).MethodByName(name).Call(nil)
}

func (t *T) GoodDinner() {
	fmt.Println("吃顿好的")
}
