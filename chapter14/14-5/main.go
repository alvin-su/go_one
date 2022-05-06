package main

import (
	"fmt"
	"reflect"
)

type IFly interface {
	Fly()
}

type Bird struct {
	name string
}

func (b *Bird) Fly() {
	fmt.Println("我会飞")
}

func main() {
	bird := &Bird{}
	t := reflect.TypeOf((*IFly)(nil)).Elem()
	refType := reflect.TypeOf(bird)
	fmt.Println(refType.Implements(t))
}
