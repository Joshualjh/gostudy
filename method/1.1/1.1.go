package main

import (
	"fmt"
)

type myInt int

func (m myInt) Add(a int) myInt {
	rst := int(m) + a
	return myInt(rst)
}

func addFunc(m myInt, a int) myInt {
	rst := int(m) + a
	return myInt(rst)
}

func main() {
	var a myInt = 10
	a = a.Add(100)
	a = addFunc(a, 100)
	fmt.Println(a)

}
