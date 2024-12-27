package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("./main.go")
	if err != nil {
		fmt.Println("打开文件失败： %v\n", err)
		return
	}
	defer file.Close()
	// 方法1
	// content, err := io.ReadAll(file)
	// if err != nil {
	// 	fmt.Println("读取文件失败：%v\n", err)
	// 	return
	// }
	// fmt.Println(string(content))
	// 方法2
	var fileStr string
	buff := make([]byte, 1024)
	for {
		n, err := file.Read(buff)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("读取文件失败：%v\n", err)
			return
		}
		fileStr += string(buff[:n])
	}

	fmt.Println(fileStr)
	// 方法3， ioutil
	// content, err := ioutil.ReadFile("./main.go")
	// if err != nil {
	// 	fmt.Println("读取文件失败：%v\n", err)
	// 	return
	// }

	// fmt.Println(string(content))

}
