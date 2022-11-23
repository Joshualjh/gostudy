package main

import (
	"fmt"
)

type Data struct {
	value int
	data  [200]int
}

func ChangeData(arg *Data) {
	arg.value = 999
	arg.data[100] = 999
}

func val(a *int) {
	*a = 100
	fmt.Println(a)
	fmt.Println(*a)
}

func main() {
	var data Data
	ChangeData(&data)
	var b int
	val(&b)
	fmt.Println(b)
	fmt.Printf("value = %d\n", data.value)
	fmt.Printf("data[100] = %d\n", data.data[100])

}
