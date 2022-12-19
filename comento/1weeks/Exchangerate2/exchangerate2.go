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

func getData(c chan *ExchangeRate, n int) error {
	for i := 0; i < n; i++ {
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
		resp.Body.Close()
		//fmt.Println(resJson.Rates["USD"])
	}

	wg.Done()
	close(c)

	return nil
}

func result(c chan *ExchangeRate, n int) error {
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
	fmt.Println(data / float64(n))
	wg.Done()
	return nil
}

func main() {
	start := time.Now()
	n := 10
	c := make(chan *ExchangeRate, n)

	wg.Add(2)
	go getData(c, n)
	//fmt.Println(i, "출력")
	// 모든 고루틴이 끝날때까지 기다림
	// 채널닫기
	go result(c, n)
	wg.Wait()
	end := time.Since(start)
	fmt.Println(end)
	// var data float64
	// for i, success := <-c; !success; {
	// 	data += float64(i.Rates["USD"])
	// 	fmt.Println(i)
	// }

	// fmt.Println(data)
}
