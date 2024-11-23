package main

import (
	"fmt"
	"go_study/demo_study/study"
)

func main() {
	name := "Tom"
	s, err := study.New(name)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(s.Listen("english"))
	fmt.Println(s.Speak("english"))
	fmt.Println(s.Read("english"))
	fmt.Println(s.Write("english"))
}
