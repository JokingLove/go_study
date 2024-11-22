package main

import "fmt"

func main() {

	demo1()
}

func demo1() {
	var a = 1
	var b = 2
	// 结论：传参是值复制。

	defer func(a, b int) {
		fmt.Println(a + b)
	}(a, b)

	a = 2
	fmt.Println("main")
}

func demo2() {
	var a = 1
	var b = 2

	//结论：闭包获取变量相当于引用传递，而非值传递。
	defer func() {
		fmt.Println(a + b)
	}()

	a = 2
	fmt.Println("main")
}
