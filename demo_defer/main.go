package main

import "fmt"

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

// 结论：defer 函数定义的顺序 与 实际执的行顺序是相反的，也就是最先声明的最后才执行。

func main() {
	x := 1
	y := 2
	defer calc("A", x, calc("B", x, y))
	x = 3
	defer calc("C", x, calc("D", x, y))
	y = 4
}
