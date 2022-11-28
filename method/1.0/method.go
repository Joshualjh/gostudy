package main

import "fmt"

type account struct {
	balance int
}

func withrawFunc(a *account, amount int) {
	a.balance -= amount
}

func (a *account) withrawMethod(amount int) {
	a.balance -= amount
}

func main() {
	a := &account{100} //*account
	withrawFunc(a, 30)
	fmt.Println(a.balance)
	a.withrawMethod(30)
	fmt.Println(a.balance)

}
