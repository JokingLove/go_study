package main

import "fmt"

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			ch1 <- i
		}
		close(ch1)
	}()
	go func() {
		for i := 0; i < 10; i++ {
			ch2 <- i
		}
		close(ch2)
	}()

	outs := merge(ch1, ch2)
	for out := range outs {
		fmt.Println("接收到的数据：", out)
	}
}

func merge(ch1, ch2 <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for {
			select {
			case v, ok := <-ch1:
				if ok {
					out <- v
				} else {
					ch1 = nil // 关闭ch1后，将其设置为  nil 之后，以避免再接收
				}
			case v, ok := <-ch2:
				if ok {
					out <- v
				} else {
					ch1 = nil // 关闭ch1后，将其设置为  nil 之后，以避免再接收
				}
			}
			if ch1 == nil && ch2 == nil {
				break
			}
		}
	}()
	return out
}
