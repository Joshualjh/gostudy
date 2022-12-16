package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup // 다른 func에도 사용할 수 있도록 전역 변수 설정
var mu sync.Mutex     // 다른 func에도 사용할 수 있도록 전역 변수 설정

func main() {
	start := time.Now() //time start
	var sum int64

	n := 5    //goroutine의 갯수
	wg.Add(n) //sync 갯수
	for i := 0; i < 10000000000; i += 10000000000 / n {
		go sumation(i, i+10000000000/n, &sum)
	}
	wg.Wait()                    //모든 goroutin이 끝날때까지 기다림 이 갯수는 위에 Add에서 지정해주고 Done의 갯수와 add의 갯수가 맞는지 확인
	elapsed := time.Since(start) // time start 부터 지금까지 걸린 시간 체크
	fmt.Println(sum, elapsed)
}

func sumation(start, end int, sumTotal *int64) {
	var sum int64
	for ; start < end; start++ {
		sum += int64(start)
	}
	mu.Lock() //다른 고루틴이 변수에 접근하지 못하도록 Lock을 걸어둠
	*sumTotal += sum
	mu.Unlock() //작업이후 Lock 해제
	wg.Done()   //한개의 고루틴이 끝나면 Done
}
