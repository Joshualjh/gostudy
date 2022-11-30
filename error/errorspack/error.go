package main

import (
	"errors"
	"fmt"
	"math"
)

func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("양수가 아니여") //fmt.Errorf("양수가 아니여 f:%g",f) error 생성 방법 두가지
	}
	return math.Sqrt(f), nil
}

func main() {
	sqrt, err := sqrt(-2)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println(sqrt)

}
