package main

import "fmt"

func main() {
	println("Hello world!")
	demo1()
	demo2()
	demo3()
	demo4()
	demo5()
	demo6()
}

func demo6() {
	nums := []int{2, 3, 4}
	sum := 0

	for _, num := range nums {
		sum += num
	}

	fmt.Println("sum:", sum)

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s \n", k, v)
	}

	for i, c := range "go" {
		fmt.Printf("%s, %d \n", c, i)
	}
}

func demo5() {
	nums := []int{2, 3, 4}

	for _, num := range nums {
		fmt.Println("value:", num)
	}

	for i := range nums {
		fmt.Println("index: ", i)
	}
}

func demo4() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	close(ch)

	for v := range ch {
		fmt.Println(v)
	}
}

func demo3() {
	map1 := make(map[int]float32)

	map1[1] = 1.0
	map1[2] = 2.0
	map1[3] = 3.0
	map1[4] = 4.0

	for key, value := range map1 {
		fmt.Printf("key is: %d - value is : %f \n", key, value)
	}

	for key, _ := range map1 {
		fmt.Printf("key is : %d \n", key)
	}

	for _, value := range map1 {
		fmt.Printf("value is: %f \n", value)
	}
}

func demo1() {
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}

func demo2() {
	for i, c := range "hello" {
		fmt.Printf("index: %d, char: %c \n", i, c)
	}
}
