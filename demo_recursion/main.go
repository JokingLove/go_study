package main

import "fmt"

func main() {
	fmt.Println("15的阶乘是", demo1(15))
}

func sqrtRecursive(x, guess, prevGuess, epsion float64) float64 {

}

func demo1(n uint64) (result uint64) {
	if n > 0 {
		result = n * demo1(n-1)
		return result
	}
	return 1
}

func fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return fibonacci(n-2) + fibonacci(n-1)
}
