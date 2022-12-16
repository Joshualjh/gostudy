package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

type ExchangeRate struct {
	Amount float64            `json:"amount"`
	Base   string             `json:"base"`
	Date   string             `json:"date"`
	Rates  map[string]float64 `json:"rates"`
}

var wg sync.WaitGroup

func getData(c chan *ExchangeRate) error {
	resp, err := http.Get("https://api.frankfurter.app/latest?from=KRW")
	if err != nil {
		return err
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	resJson := ExchangeRate{}
	if err := json.Unmarshal(data, &resJson); err != nil {
		return err
	}

	c <- &resJson
	resp.Body.Close()
	wg.Done()
	return nil
}

func result(c chan *ExchangeRate) error {
	var data float64
	// for i := range c {
	// 	data += i.Rates["USD"]
	// 	println(data)
	// }

	for {
		if i, success := <-c; success {
			data += i.Rates["USD"]
			fmt.Println(data)
		} else {
			break
		}
	}
	fmt.Println(data / 10)

	return nil
}

func main() {
	c := make(chan *ExchangeRate, 10)
	n := 10
	for i := 0; i < n; i++ {
		wg.Add(1)
		go getData(c)
		fmt.Println(i, "출력")
	}
	wg.Wait()
	close(c)
	result(c)

	// var data float32
	// for i, success := <-c; !success; {
	// 	data += float32(i.Rates["USD"])
	// 	fmt.Println(i)
	// }

	// fmt.Println(data)
}
