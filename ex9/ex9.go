package main

import (
	"fmt"
)

func rich(a bool) bool {
	return a
}

func setprice(p int, v int) {
	fmt.Println(v / p)
}

func main() {
	var value int
	var person int
	var rich bool
	fmt.Println("음식 값이 얼마입니까?")
	fmt.Scanln(&value)
	fmt.Println("몇명입니까?")
	fmt.Scanln(&person)
	fmt.Println("부자가 있습니까?")
	fmt.Scanln(&rich)
	if value > 50000 {
		if rich {
			fmt.Println("신발끈을 묶는다")
		} else {
			setprice(person, value)
		}
	} else if value >= 30000 {
		if person > 3 {
			fmt.Println("신발끈을 묶는다")
		} else {
			setprice(person, value)
		}
	} else {
		fmt.Println("내가 낸다")
	}

}
