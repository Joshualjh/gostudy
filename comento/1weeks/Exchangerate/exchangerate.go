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
	defer wg.Done()
	resp, err := http.Get("https://api.frankfurter.app/latest?from=KRW")
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	resJson := ExchangeRate{}
	if err := json.Unmarshal(data, &resJson); err != nil {
		return err
	}

	c <- &resJson

	return nil
}

func result(c chan *ExchangeRate) error {
	var data float64
	for {
		if i, success := <-c; success {
			println(i)
		} else {
			break
		}
	}

	fmt.Println(data)
	return nil
}

func main() {
	c := make(chan *ExchangeRate)
	n := 5
	for i := 0; i < n; i++ {
		wg.Add(1)
		go getData(c)
		fmt.Println(i, "출력")
	}

	result(c)
	close(c)
	wg.Wait()
	// var data float32
	// for i, success := <-c; !success; {
	// 	data += float32(i.Rates["USD"])
	// 	fmt.Println(i)
	// }

	// fmt.Println(data)
}
