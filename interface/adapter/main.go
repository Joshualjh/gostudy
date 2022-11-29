package main

import "fmt"

type Database interface {
	Get()
	Set()
}

type CDatabase struct {
}

func (c CDatabase) GetData() {
	fmt.Print("test")
}

func (c CDatabase) SetData() {
}

// 서로 다른 메서드 끼리 interface에서 적용하기 위해 한번 감싸주는 역활
type Wrapper struct {
	cdb CDatabase
}

func (w Wrapper) Get() {
	w.cdb.GetData()
}

func (w Wrapper) Set() {
	w.cdb.SetData()
}

func main() {
	var cdatabase CDatabase
	var database Database
	database = Wrapper{cdatabase}
	database.Get()

}
