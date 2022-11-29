package main

import (
	"fmt"
	"os"
)

type Writer func(string)
type Writerinterface interface {
	Write(string)
}

func writeHello(writer Writer) {
	writer("Hello World")
} //의존성 주입 방식 - 외부에서 어떤 내용을 주입하는 방식

func main() {
	f, err := os.Create("test.txt")
	if err != nil {
		fmt.Println("fail")
		return
	}
	defer f.Close()

	writeHello(func(msg string) {
		fmt.Println(msg)
	})
}
