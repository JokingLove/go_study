package main

import "fmt"

func main() {
	demo1()
	demo2()
}

func demo2() {
	var a int = 4
	var b int32
	var c float32
	var ptr *int

	fmt.Printf("第一行 - a 变量类型为 = %T\n", a)
	fmt.Printf("第二行 - b 变量类型为 = %T\n", b)
	fmt.Printf("第三行 - c 变量类型为 = %T\n", c)

	ptr = &a
	fmt.Printf("a 的值为 %d\n", a)
	fmt.Printf("*ptr 为 %d\n", *ptr)
}
func demo1() {
	var a int = 21
	var b int = 10
	var c int

	c = a + b
	fmt.Println("第一行 - c 的值为 %d \n", c)

	c = a - b
	fmt.Println("第二行 - c 的值为 %d \n", c)

	c = a * b
	fmt.Println("第三行 - c 的值为 %d \n", c)

	c = a / b
	fmt.Println("第四行 - c 的值为 %d \n", c)

	c = a % b
	fmt.Println("第五行 - c 的值为 %d \n", c)

	a++
	fmt.Println("第六行 - a 的值为 %d \n", a)

	a = 21
	a--
	fmt.Println("第七行 - a 的值为 %d \n", a)
}
