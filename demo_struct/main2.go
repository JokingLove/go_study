package main

import "fmt"

type Student struct {
	Name  string
	Sex   string
	Grade string
	Age   uint
}

func NewStudent(Name string, Sex string, Grade string, Age uint) (*Student, error) {
	return &Student{
		Name:  Name,
		Sex:   Sex,
		Grade: Grade,
		Age:   Age,
	}, nil
}

func (s *Student) GetName() string {
	return s.Name
}

func (s *Student) GetSex() string {
	return s.Sex
}

func main() {
	st, err := NewStudent("haha", "男", "一年级", 18)
	if err != nil {
		fmt.Println("错误")
		return
	}

	fmt.Println(st)
}
