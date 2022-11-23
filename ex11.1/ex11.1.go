package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func NewUser(name string, age int) *User {
	var u = User{name, age}
	return &u

}

func main() {
	userPointer := NewUser("AAA", 23)
	fmt.Println(*userPointer)

}

// dangling err - 이미 사라진 주소를 반환한다.
// go에서는 Excape Analyzing 탈출분석을 하기 때문에 자료가 남아있음
