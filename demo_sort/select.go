package main

import "fmt"

// 选择排序
func main() {
	//
	arr := []int{3, 4, 2, 1, 0, 8, 5}
	// for i := 0; i < len(arr); i++ {
	// 	for j := 0; j < len(arr); j++ {
	// 		if arr[i] < arr[j] {
	// 			temp := arr[i]
	// 			arr[i] = arr[j]
	// 			arr[j] = temp
	// 		}
	// 	}
	// }

	// 冒泡排序
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				temp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
			}
		}
	}
	fmt.Println(arr)
	fmt.Printf("%d", len(arr))
}
