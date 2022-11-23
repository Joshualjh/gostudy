package main

import (
	"fmt"
	"unsafe"
)

type Student struct {
	Age    int8
	Name   string
	mobile string
}

func main() {
	var student Student
	student.Name = "이정현"
	student.Age = 27
	student.mobile = "010-1001-1001"

	fmt.Println("이름 :", student.Name, "나이 :", student.Age, "휴대폰 :", student.mobile)
	fmt.Println(unsafe.Sizeof(student))
}
