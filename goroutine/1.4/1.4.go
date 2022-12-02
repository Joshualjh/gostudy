// Deadlock 상황 만들기
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	rand.Seed(time.Now().UnixNano())

	wg.Add(2)
	fork := &sync.Mutex{}
	spoon := &sync.Mutex{}

	go diningProblem("A", spoon, fork, "수저", "포크")
	go diningProblem("B", spoon, fork, "수저", "포크")
	go diningProblem("C", spoon, fork, "수저", "포크")
	go diningProblem("D", spoon, fork, "수저", "포크")
	go diningProblem("E", spoon, fork, "수저", "포크")
	go diningProblem("F", spoon, fork, "수저", "포크")

	wg.Wait()
}

func diningProblem(name string, first, second *sync.Mutex, firstName, secondName string) {
	cnt := 0
	for i := 0; i < 100; i++ {
		fmt.Printf("%s 밥을 먹으려 합니다\n", name)
		first.Lock()
		fmt.Printf("%s %s 획득\n", name, firstName)
		second.Lock()
		fmt.Printf("%s %s 획득\n", name, secondName)
		fmt.Printf("%s 밥을 먹습니다\n", name)
		cnt += 1
		fmt.Printf("%d 번 밥을 먹었습니다.\n", cnt)
		time.Sleep(time.Duration(rand.Intn(1)) * time.Millisecond)
		second.Unlock()
		first.Unlock()
	}
	fmt.Printf("%s가 밥을 먹은 횟수 %d 번입니다.", name, cnt)
	wg.Done()
}
