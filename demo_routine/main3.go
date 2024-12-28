package main

import (
	"fmt"
	"time"
)

func main() {
	// sele1()

	// select2()

	// select3()

	select4()
}

func select4() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		ch1 <- "channel 1"
		ch2 <- "channel 2"
	}()

	// ch1 具有更高优先级
	for i := 0; i < 20; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println("高优先级：", msg1)
		default:
			select {
			case msg2 := <-ch2:
				fmt.Println("低优先级：", msg2)
			default:
				fmt.Println("都没有准备好")
			}
		}
	}
}

func select3() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	done := make(chan bool)

	// 生产者
	go func() {
		for i := 0; i < 5; i++ {
			// time.Sleep(time.Second)
			ch1 <- fmt.Sprintf("ch1: %d", i)
		}
	}()

	go func() {
		for i := 0; i < 5; i++ {
			// time.Sleep(2 * time.Second)
			ch2 <- fmt.Sprintf("ch2: %d", i)
		}
		done <- true
	}()

	// 消费者
	for {
		select {
		case msg1 := <-ch1:
			fmt.Println(msg1)
		case msg2 := <-ch2:
			fmt.Println(msg2)
		case <-done:
			fmt.Println("全部完成")
			return
		}
	}

}

func select2() {
	ch := make(chan string)
	go func() {
		time.Sleep(2 * time.Second)
		ch <- "处理完成"
	}()

	select {
	case result := <-ch:
		fmt.Println(result)
	case <-time.After(3 * time.Second):
		fmt.Println("超时了！")
	}
}

func sele1() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- "channel 1"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		ch2 <- "channel 2"
	}()

	for i := 0; i < 22; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println(msg1)
		case msg2 := <-ch2:
			fmt.Println(msg2)
		}
	}
}
