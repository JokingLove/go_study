package main

import "fmt"

/**
 Channel 分为有缓冲区和无缓冲区
 无缓冲区 channel ： 接收者和发送者必须同时存在，发送和接收都会阻塞
 缓冲区 channel： 允许再没有接受者的情况下继续发送数据，直到达到缓冲区的大小，达到缓冲区最大承载值，继续发送消息会阻塞
**/

// 无缓冲区
func fn1() {
	ch := make(chan int)
	go func() {
		ch <- 42
	}()

	value := <-ch
	fmt.Println("接收到的值为：", value)
}

// 有缓冲区的channel
func fn2() {
	tasks := make(chan int, 10)
	for i := 0; i < 10; i++ {
		fmt.Println("i===", i)
		tasks <- i
	}
	close(tasks)

	for task := range tasks {
		fmt.Println("task == ", task)
	}
}

func fn3() {
	ch0 := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			ch0 <- i
		}
		close(ch0)
	}()

	for v := range ch0 {
		fmt.Println("接收消息=", v)
	}

	fmt.Println("channel 结束")
}

func fn4() {
	ch0 := make(chan int, 5)
	go func() {
		for i := 0; i < 5; i++ {
			ch0 <- i
		}
		close(ch0)
	}()

	for v := range ch0 {
		fmt.Println("接收消息=", v)
	}

	fmt.Println("channel 结束")
}

func main() {
	// fn1()
	// fn2()
	fn3()
}
