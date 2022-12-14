package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var stdin = bufio.NewReader(os.Stdin)

func create() [3]int {
	rand.Seed(time.Now().UnixMicro())
	r := [3]int{}
	for r[0] = rand.Intn(10); r[0] == 0; r[0] = rand.Intn(10) {
	}
	for r[1] = rand.Intn(10); r[0] == r[1] || r[1] == 0; r[1] = rand.Intn(10) {
	}
	for r[2] = rand.Intn(10); r[2] == r[1] || r[2] == r[0] || r[2] == 0; r[2] = rand.Intn(10) {
	}
	fmt.Println(r)
	return r
}

//func inputValue() {
//	var i [3]int
//}

func main() {
	fmt.Println("start number baseball")
	fmt.Println("make random number")
	var r [3]int = create()
	v := "123"
	fmt.Println(strconv.F(v[2]))
	fmt.Println(r)
}
