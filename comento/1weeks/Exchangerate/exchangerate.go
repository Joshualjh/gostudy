package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

type ExchangeRate struct {
	Amount float64            `json:"amount"`
	Base   string             `json:"base"`
	Date   string             `json:"date"`
	Rates  map[string]float64 `json:"rates"`
}

var wg sync.WaitGroup

func getData(c chan *ExchangeRate) error {
	resp, err := http.Get("https://api.frankfurter.app/latest?from=KRW") //환율정보 가져옴
	if err != nil {
		return err
	}
	//fmt.Println(ioutil.ReadAll(resp.Body))
	//fmt.Println(resp.Body)
	data, err := ioutil.ReadAll(resp.Body) // EOF까지 resp의 body 값을 읽어서 저장 저장형식은 byte
	if err != nil {                        // err 처리
		return err
	}
	//fmt.Println(data)
	resJson := ExchangeRate{}                              //형식
	if err := json.Unmarshal(data, &resJson); err != nil { //byte 값으로 된 데이터를 시스템 처리 가능한데이터로 언마샬링
		return err
	}

	c <- &resJson //채널에 저장
	defer resp.Body.Close()
	defer wg.Done()
	return nil
}

func result(c chan *ExchangeRate) error {
	var data float64
	for i := range c {
		data += i.Rates["USD"]
		//fmt.Println(data)
	}

	// for {
	// 	if i, success := <-c; success {
	// 		data += i.Rates["USD"]
	// 		fmt.Println(data)
	// 	} else {
	// 		break
	// 	}
	// }
	fmt.Println(data / 10)

	return nil
}

func main() {
	start := time.Now()
	c := make(chan *ExchangeRate, 10)
	n := 10
	for i := 0; i < n; i++ {
		wg.Add(1)
		go getData(c)
		//fmt.Println(i, "출력")
	}
	wg.Wait() // 모든 고루틴이 끝날때까지 기다림
	close(c)  // 채널닫기
	result(c)
	end := time.Since(start)
	fmt.Println(end)
	// var data float32
	// for i, success := <-c; !success; {
	// 	data += float32(i.Rates["USD"])
	// 	fmt.Println(i)
	// }

	// fmt.Println(data)
}
