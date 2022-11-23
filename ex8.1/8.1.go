package main

import "fmt"

func main() {
	const (
		BitFlag1 uint = 1 << iota
		BitFlag2
		BitFlag3
		BitFlag4
	)
	fmt.Println(BitFlag1, BitFlag2, BitFlag3, BitFlag4)
}
