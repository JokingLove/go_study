package main

import "fmt"

func main() {
	var a int = 10

	// fmt.Printf("变量的地址： %x \n", &a)
	// demo2()

	// println("-=--------")
	// demo3()
	var p = &a

	fmt.Println(fn1(p))
}

func fn1(*a int) int {
	return a
}

func demo2() {
	var a int = 20
	var ip *int

	ip = &a
	fmt.Printf("a 变量的地址是： %x \n", &a)

	fmt.Printf("ip 变量存储的指针地址： %x \n", ip)

	fmt.Printf("*ip 变量的值： %d \n", *ip)
}

func demo3() {
	var ptr *int

	fmt.Printf("ptr 的值为： %x\n", ptr)

	if ptr == nil {
		fmt.Printf("ptr 的值为： nil\n")
	}
}
