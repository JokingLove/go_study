package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name   string `json:"name" validate:"required"`
	Age    int    `json:"age"1`
	Gender string
}

func (p Person) SayHello(msg string) string {
	return fmt.Sprintf("Hello, %s. I am %s", msg, p.Name)
}

func main() {
	p := Person{
		Name:   "张三",
		Age:    25,
		Gender: "男",
	}
	// 1.获取类型信息
	t := reflect.TypeOf(p)
	fmt.Printf("类型名称：%v\n", t.Name())
	fmt.Printf("类型种类：%v\n", t.Kind())
	fmt.Printf("类型字符串：%v\n", t.String())

	// 2.获取值信息
	v := reflect.ValueOf(p)
	fmt.Printf("值：%v\n", v)

	fmt.Println("---------------")
	// 3.遍历结构体字段
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)
		// 获取字段标签
		jsonTag := field.Tag.Get("json")
		fmt.Printf("字段名：%s, 类型：%s，值：%v, json标签：%s\n", field.Name, field.Type, value, jsonTag)
	}

	fmt.Println("---------------")

	// 4.修改值（需要传入指针）
	pv := reflect.ValueOf(&p).Elem()
	if nameField := pv.FieldByName("Name"); nameField.CanSet() {
		nameField.SetString("李四")
	}
	fmt.Printf("修改后的名字：%s\n", p.Name)
	fmt.Println("---------------")

	// 5.调用方法
	// method := p.MethodByName("SayHello") // 方法调用结果：Hello, 世界. I am 张三
	method := pv.MethodByName("SayHello") // 方法调用结果：Hello, 世界. I am 李四
	if method.IsValid() {
		args := []reflect.Value{reflect.ValueOf("世界")}
		result := method.Call(args)
		fmt.Printf("方法调用结果：%v\n", result[0].String())
	}

	fmt.Println("---------------")
	// 6. 创建新的实例
	newStruct := reflect.New(t).Elem()
	newStruct.FieldByName("Name").SetString("王五")
	newStruct.FieldByName("Age").SetInt(30)
	fmt.Println("newStruct", newStruct)
	newPerson := newStruct.Interface().(Person)
	fmt.Printf("新创建的实例：%+v\n", newPerson)
}
