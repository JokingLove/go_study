package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("main start")

	ch := make(chan string, 1)
	ch <- "a"
	go func() {
		val := <-ch
		fmt.Println(val)
	}()
	fmt.Println("main end")
	time.Sleep(1 * time.Second)
}
