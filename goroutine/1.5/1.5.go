package main

import (
	"fmt"
	"sync"
	"time"
)

type job interface {
	Do()
}

type SquareJob struct {
	index int
}

func (j *SquareJob) Do() {
	fmt.Printf("%d 작업 시작 \n", j.index)
	time.Sleep(1 * time.Second)
	fmt.Printf("%d 작업완료 - 결과 %d\n", j.index, j.index*j.index)
}

func main() {
	var joblist [10]job
	for i := 0; i < 10; i++ {
		joblist[i] = &SquareJob{i}
	}
	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		job := joblist[i]
		go func() {
			job.Do()
			wg.Done()
		}()
	}
	wg.Wait()
}
