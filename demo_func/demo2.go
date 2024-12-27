package main

import "fmt"

func main() {
	fmt.Println(fn1(10, 0))
	fmt.Println("end")
}

func fn1(a, b int) int {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("err: ", r)
			return
		}
	}()
	return a / b
}
