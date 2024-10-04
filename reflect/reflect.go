package main

import (
	"fmt"
	"reflect"
)

type Address struct {
	City  string
	Phone string
}

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`

	// 嵌套结构体
	Address Address `json:"address"`
}

func main() {
	// 创建一个 Person 实例
	p := Person{
		Name: "Alice",
		Age:  30,
		Address: Address{
			City:  "Beijing",
			Phone: "123456789",
		},
	}

	// 获取 Person 的类型信息
	t := reflect.TypeOf(p)
	fmt.Println("Type:", t)

	// 获取 Person 的字段信息
	v := reflect.ValueOf(p)

	fmt.Println("Value:", v)

	// 如何获取结构体所有字段对应的值?
	for i := 0; i < v.NumField(); i++ { // 需要变量对象值信息
		field := v.Field(i)
		fmt.Printf("Field %d: %v\n", i, field, field.FieldByName("Name"))
	}

	// 如何获取结果体中字段名和对应类型？
	for i := 0; i < t.NumField(); i++ { // 需要遍历类型信息
		field := t.Field(i)
		fmt.Printf("Field %d:  %v %v %v\n", i, field.Name, field.Type, field.Tag)
	}

}
