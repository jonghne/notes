package main

import (
	"fmt"
	"reflect"
)

type frac struct {}

type fu interface {
	Do(in string)
}

func (f *frac) Do(in string) {
	fmt.Println(in)
}

func mao(f fu) {
	fmt.Println(reflect.TypeOf(f), reflect.ValueOf(f), reflect.ValueOf(f).NumMethod())
}

func main() {
	dev := frac{}
	fmt.Println(reflect.TypeOf(&dev), reflect.ValueOf(&dev), reflect.ValueOf(&dev).NumMethod())
	mao(&dev)
	(&dev).Do("haha")
	fr := reflect.ValueOf(&dev).MethodByName("Do")
	fmt.Println(fr)
	if !fr.IsValid() || fr.IsNil() {
		fmt.Println("panic")
		return
	}

	fr.Call([]reflect.Value{reflect.ValueOf("jicai")})
}