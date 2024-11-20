package main

import (
	"fmt"
	"unsafe"
)

func main() {
	demo1()
	demo2()
	demo3()

}

func demo3() {
	const (
		a = iota
		b
		c
		d = "ha"
		e
		f = 100
		g
		h = iota
		i
	)
	fmt.Println(a, b, c, d, e, f, g, h, i)
}

const (
	a = "abc"
	b = len(a)
	c = unsafe.Sizeof(a)
)

func demo2() {
	println(a, b, c)
}

func demo1() {
	const LENGTH int = 10
	const WIDTH int = 5
	var area int
	const a, b, c = 1, false, "str"

	area = LENGTH * WIDTH
	fmt.Printf("面积为： %d", area)
	println()
	println(a, b, c)
}
