package main

import (
	"fmt"
	"sync"
)

var ch = make(chan int, 20)
var wg sync.WaitGroup

func main() {
	wg.Add(1)
	go putVal()
	wg.Add(1)
	go printVal()

	wg.Wait()
	// close(ch)
}

func putVal() {
	for i := 0; i < 100; i++ {
		ch <- i
	}
	close(ch)
	wg.Done()
}

func printVal() {
	defer wg.Done()
	for num := range ch {
		fmt.Println(num)
	}
}
