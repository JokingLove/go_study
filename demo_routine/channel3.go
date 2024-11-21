package main

import "fmt"

// 并发编程小结
// Goroutines 是轻量级线程，使用 go 关键字启动。
// Channels 用于 goroutines 之间的通信。
// Select 语句 用于等待多个 channel 操作。

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
	fmt.Println("hello world!")
}
