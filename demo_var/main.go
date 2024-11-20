package main

import "fmt"

func main() {
	var a string = "Runoob"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)
	demo1()
	demo2()
	demo3()
	demo4()
	demo5()
}

func demo5() {
	_, numb, strs := numbers()
	fmt.Println(numb, strs)
}
func numbers() (int, int, string) {
	a, b, c := 1, 2, "str"
	return a, b, c
}
func demo1() {
	var a = "RUNOOB"
	fmt.Println(a)

	var b int
	fmt.Println(b)

	var c bool
	fmt.Println(c)
}

func demo2() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q \n", i, f, b, s)
}

func demo3() {
	f := "RUNOOB"

	fmt.Println(f)
}

var x, y int
var (
	a int
	b bool
)

var c, d int = 1, 2
var e, f = 123, "hello"

func demo4() {
	g, h := 123, "world"
	fmt.Println(x, y, a, b, c, d, e, f, g, h)
}
