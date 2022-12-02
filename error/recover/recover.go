package main

import "fmt"

func f() {
	fmt.Println("f() 함수시작")
	defer func() {
		if r := recover(); r != nil { //panic에 대한 오류를 DB에 저장한 후에 panic을 해결하는 것이 더 좋습니다..
			fmt.Println("panic 복구 - ", r)
		}
	}()

	g()
	fmt.Println("f() 함수 끝")
}

func g() {
	fmt.Printf("9/3 = %d\n", h(9, 3))
	fmt.Printf("9/0 = %d\n", h(9, 0))

}
func h(a, b int) int {
	if b == 0 {
		panic("제수는 0일수 없습니다.")
	}
	return a / b
}

func main() {
	f()
	fmt.Println("프로그램이 계속 실행됨")
}
