package main

import (
	"fmt"
)

type NameBook struct {
	Address string
	age     int
}

func main() {
	m := make(map[string]NameBook)
	m["이정현"] = NameBook{"경기도 양평군", 27}
	m["홍길동"] = NameBook{"서울시 강남구", 17}
	m["샤오미"] = NameBook{"부천시", 25}
	m["니달리"] = NameBook{"서울시 강서구", 50}

	println(m)

	for i, v := range m {
		fmt.Printf("%s는 %d살이며 %s에 삽니다.\n", i, v.age, v.Address)
	}

}
