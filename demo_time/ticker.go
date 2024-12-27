package main

import (
	"fmt"
	"time"
)

func main() {
	fn1()
}

func fn1() {
	tiker := time.NewTicker(time.Second)
	n := 5
	for t := range tiker.C {
		n--
		fmt.Println(t)
		if n == 0 {
			tiker.Stop()
			break
		}
	}
}
