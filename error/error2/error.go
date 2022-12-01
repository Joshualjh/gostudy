package main

import (
	"bufio"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func MultipleFromString(str string) (int, error) {
	//데이터를 가져올때 일정한 규칙으로 가져옴 인자로 IOReader를 받음
	scanner := bufio.NewScanner(strings.NewReader(str))
	scanner.Split(bufio.ScanWords)

	pos := 0
	a, n, err := readNextInt(scanner)
	if err != nil {
		return 0, fmt.Errorf("Failed to readNextInt(), pod:%d err:%w", pos, err)
	}

	pos += n + 1
	b, n, err := readNextInt(scanner)
	if err != nil {
		return 0, fmt.Errorf("Failed to readNextInt(), pod:%d err:%w", pos, err)
	}
	return a * b, nil

}

// 다음 단어를 읽어서 숫자로 변환하여 반환
// 변환된 숫자, 읽은 글자수, 에러를 반환
func readNextInt(scanner *bufio.Scanner) (int, int, error) {
	if !scanner.Scan() {
		return 0, 0, fmt.Errorf("Failed to scan")
	}
	word := scanner.Text()
	number, err := strconv.Atoi(word) //"24" -> 24 , "abc"-> error
	if err != nil {
		return 0, 0, fmt.Errorf("Failed to convert word to int word:%s err:%w", word, err)
	}
	return number, len(word), nil

}

func readEq(eq string) {
	rst, err := MultipleFromString(eq)
	if err == nil {
		fmt.Println(rst)
	} else {
		fmt.Println(err)
		var numError *strconv.NumError
		if errors.As(err, &numError) { //error를 numError로 감싸서 넘기는 방법
			fmt.Println("NumberError", numError)
		}
	}
}

func main() {
	readEq("123 3")
	readEq("123 abc")
}
