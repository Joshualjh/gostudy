package main

import "fmt"

func print[T any](t T) {
	fmt.Println(t)
}

func main() {
	print(1)
	print("test")
	print(3.14)
	print(true)
}
