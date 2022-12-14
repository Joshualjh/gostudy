package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a int // 변수선언
	var f float32 = 11.0
	f = 12.0
	const b int = 13 // 상수선언
	//b = 23 // 상수이므로 변경불가
	fmt.Println(a, f)
	fmt.Println(reflect.TypeOf(a), reflect.TypeOf(f))

	var p = &a         //a의 주소값 할당
	fmt.Println(p, *p) // (주소, 값)출력
}
