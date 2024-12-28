package main

import (
	"fmt"
	"sync"
)

var (
	i    = 0
	lock sync.Mutex
	wg   sync.WaitGroup
)

func main() {
	for j := 0; j < 20; j++ {
		wg.Add(1)
		go printInfo()
	}

	wg.Wait()
}

func printInfo() {
	lock.Lock()
	fmt.Println(i)
	i++
	lock.Unlock()
	wg.Done()
}
