package main

import (
	"fmt"
	"reflect"
	"strings"
)

type boy struct {
	Name string `json:"sign_number" gorm:"type:varchar(255);column:sign_number" fname:"标记编号"`
	age  int
}

type human interface {
	SayName(p string)
	SayAge()
}

func (this *boy) SayName(p string) {
	fmt.Println(this.Name, "call", p)
}

func (this *boy) SayAge() {
	fmt.Println(this.age)
}

func main() {
	// 定义接口变量
	var i human
	// 初始化对象，jown持有对象指针。
	jown := &boy{
		Name: "jown",
		age: 15,
	}
	// 因为boy实现了human中的方法，所以它实现了human接口。
	// 这时，i就指向jown对象。
	i = jown

	// 通过反射获取接口i 的类型和所持有的值。
	t := reflect.TypeOf(i)
	v := reflect.ValueOf(i)

	fmt.Println(t, v)
	fmt.Println(t.Elem(), t.Elem().Name(), t.Elem().Field(0), t.Elem().Field(1))
	fmt.Println(t.Elem().Field(0).Tag.Get("json"), t.Elem().Field(0).Tag.Get("gorm"))
	fmt.Println(strings.Split(t.Elem().Field(0).Tag.Get("gorm"), ";"))
	fmt.Println(v.Elem(), v.Elem().NumField())

	// 获取i所指向的对象的类型
	structType := t.Elem()
	// 获取对象的名字
	structName := structType.Name()
	fmt.Println(structType, structName)
	fmt.Println(reflect.New(t.Elem()).NumMethod(), reflect.New(t.Elem()).Method(1).Type().In(0))
	fmt.Println(reflect.TypeOf(reflect.New(t.Elem()).Interface()))

	method, _ := t.MethodByName("SayAge")
	fmt.Println(method)

	fmt.Println(v.Elem())
	vmethod := v.MethodByName("SayName")
	fmt.Println(vmethod, vmethod.Type().NumIn())
	args := []reflect.Value{reflect.ValueOf("liming")}
	// 通过v进行调用
	v.MethodByName("SayName").Call(args)
}
