package main

import (
	"fmt"
	"strconv"
)

func main() {
	demo1()
	demo2()
	demo3()
	demo4()
	demo5()

	printValue(432)
	printValue("Hellorr")
	printValue(3.14)
}

func printValue(v interface{}) {
	switch v := v.(type) {
	case int:
		fmt.Println("Integer:", v)
	case string:
		fmt.Println("String:", v)
	default:
		fmt.Println("Unknown type")
	}
}

type Writer interface {
	Write([]byte) (int, error)
}

type StringWriter struct {
	str string
}

func (sw *StringWriter) Write(data []byte) (int, error) {
	sw.str += string(data)
	return len(data), nil
}

func demo5() {
	var w Writer = &StringWriter{}
	sw := w.(*StringWriter)

	sw.str = "Hello, World"

	fmt.Println(sw.str)

	len, err := sw.Write([]byte{1, 2, 3})
	fmt.Println(len)
	fmt.Println(err)
	fmt.Println(sw.str)
}

// 接口类型断言
func demo4() {
	var i interface{} = "Hello, World!"
	str, ok := i.(string)
	if ok {
		fmt.Printf("'%s' is a string \n", str)
	} else {
		fmt.Println("conversion failed")
	}
}

func demo3() {
	str := "3.14"
	num, err := strconv.ParseFloat(str, 64)
	if err != nil {
		fmt.Println("转换错误：", err)
	} else {
		fmt.Printf("字符串 '%s' 转为浮点型为： %f \n", str, num)
	}
}

func demo2() {
	str := "123"
	num, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("转换错误：", err)
	} else {
		fmt.Printf("字符串 '%s' 转换为整数为： %d \n", str, num)
	}
}

func demo1() {
	var sum int = 17
	count := 5
	var mean float32

	mean = float32(sum) / float32(count)
	fmt.Printf("mean 的值为： %f \n", mean)
	fmt.Println(sum / count)
}
