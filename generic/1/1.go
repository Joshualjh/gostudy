package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// 타입제한자
type Integer interface {
	int | int8 | int16 | int32 | int64
}
type Float interface {
	float32 | float64
}

// go에서 정해준 패키지를 이용한 방법
func min[T constraints.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

// generic으로 비교연산시 대소비교가 가능한 타입만으로 적어야 함.
// any의 경우 비교가 불가능한 것들이 있기 떄문에 불가능
// 타입 제한자를 정의해서 쓰는 경우가 대부분입니다.
// func min[T int | int16 | int32 | float32 | float64 | int64](a, b T) T {
// 	if a < b {
// 		return a
// 	}
// 	return b
// }
// func min[T Integer | Float](a, b T) T {
// 	if a < b {
// 		return a
// 	}
// 	return b
// }

func main() {
	var a int = 10
	var b int = 20
	fmt.Println(min(a, b))

	var c float32 = 10.4999
	var d float32 = 20.4
	fmt.Println(min(c, d))

	var e int64 = 10
	var f int64 = 20
	fmt.Println(min(e, f))
}
