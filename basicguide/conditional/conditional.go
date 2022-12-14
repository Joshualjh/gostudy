package main

import (
	"fmt"
)

func main() {
	var k = 1
	if k == 1 {
		fmt.Println(k)
	} else if k == 2 {
		fmt.Println("test")
	} else {
		fmt.Println("err")
	}

	switch k {
	case 1:
		fmt.Println("1man")
	case 2:
		fmt.Println("1woman")
	case 3:
		fmt.Println("1?")
	default:
		break
	}

	switch {
	case k == 1:
		fmt.Println("2man")
	case k == 2:
		fmt.Println("2woman")
	case k == 3:
		fmt.Println("2?")
	default:
		return

	}

	switch {
	case k == 1:
		fmt.Println("3man")
		fallthrough //아래의 케이스도 실행시킴
	case k == 2:
		fmt.Println("3woman")
		fallthrough
	case k == 3:
		fmt.Println("3?")
		fallthrough
	default:
		return

	}

}
