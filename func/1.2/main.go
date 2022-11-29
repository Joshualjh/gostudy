package main

import "fmt"

type OpFn func(int, int) int

func add(a, b int) int {
	return a + b
}
func mul(a, b int) int {
	return a * b
}

func getOperator(op string) OpFn {
	if op == "+" {
		return add
	} else if op == "*" {
		return mul
	} else {
		return nil
	}
}

func main() {
	var operator OpFn
	operator = getOperator("*")
	var result = operator(3, 4)
	fmt.Println(result)
	operator = getOperator("+")
	result = operator(3, 4)
	fmt.Println(result)
	operator = getOperator("-")
	result = operator(3, 4)
	fmt.Println(result)
}
