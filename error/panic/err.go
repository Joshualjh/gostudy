package main

import "fmt"

func divide(a, b int) {
	if b == 0 {
		//에러가 발생했을경우 프로그램을 죽이는 방법
		//문제가 어디서 나왔는지 바로 볼 수 있음
		//프로그램 강제 종료됨
		panic("b는 0일수 없음")
	}
	fmt.Printf("%d / %d = %d\n", a, b, a/b)
}

func main() {
	divide(9, 3)
	divide(10, 0)
	divide(11, 2)

}
