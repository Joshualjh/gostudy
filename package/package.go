package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	fmt.Println(rand.Int())
	fmt.Println(math.Pi)
	var r float64 = 20
	var val float64 = math.Pi * r * r
	fmt.Println(val)

}
