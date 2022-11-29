package main

import (
	"fmt"
)

func main() {
	m := make(map[int]int)
	m[1] = 1
	m[2] = 2
	m[3] = 3
	m[4] = 4
	m[5] = 5

	delete(m, 3)
	delete(m, 6)
	delete(m, 2)

	for i := 0; i < 6; i++ {
		v, ok := m[i] //v : 값 ok : 존재 여부
		fmt.Println(i, v, ok)
	}

}

/*
맵을 사용하는 것이 속도 면에서 가장 빠릅니다.
하지만 메모리를 더 씁니다.
*/
