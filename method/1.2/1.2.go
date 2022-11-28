package main

import (
	"fmt"
)

type account struct {
	balance   int
	firstname string
	lastname  string
}

func (a1 *account) withdrawPointer(amount int) {
	a1.balance -= amount
}

func (a2 account) withdrawValue(amount int) {
	a2.balance -= amount
}

func (a2 account) withdrawValue2(amount int) account {
	a2.balance -= amount
	return a2
}

func main() {
	var mainA *account = &account{100, "joe", "park"}
	//포인터 타입
	mainA.withdrawPointer(30)
	fmt.Println(mainA.balance)

	mainA.withdrawValue(30)
	fmt.Println(mainA.balance)

	//값타입
	*mainA = mainA.withdrawValue2(30)
	fmt.Println(mainA.balance)
}
