// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println(t.UnixNano())
	rand.Seed(t.UnixNano())
	for i := 0; i < 10; i++ {
		fmt.Println(rand.Intn(100))
	}
}
