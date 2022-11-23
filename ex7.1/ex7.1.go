package main

import (
	"fmt"
)

func Add(a int, b int) int {
	return a + b
}

func main() {
	var a int
	var b int
	n, err := fmt.Scanln(&a, &b)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(n)
	}
	c := Add(a, b)
	fmt.Println(c)
}
