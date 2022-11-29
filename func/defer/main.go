package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("test.txt")
	if err != nil {
		fmt.Println("Failed to create a file", err)
		return
	}
	defer fmt.Println("반드시 호출됨")
	defer f.Close()
	defer fmt.Println("파일 닫기")
	fmt.Println("파일에 Hello World쓰기")
	fmt.Fprintln(f, "Hello World")

}
