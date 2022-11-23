package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6}
	idx := 2

	//슬라이스 중간에 값을 없앨때 할때 사용하는 방법
	//---------------------------------
	// slice = append(slice[:idx], slice[idx+1:]...)
	//--
	// for i := idx + 1; i < len(slice); i++ {
	// 	slice[i-1] = slice[i]
	// }
	// slice = slice[:len(slice)-1]
	//---------------------------------
	//슬라이스 중간 값을 넣을때 ex[1,2,3,100,4,5,6]
	// slice = append(slice[:idx], append([]int{100}, slice[idx:]...)...)
	//--
	// slice = append(slice, 0)
	// for i := len(slice) - 2; i >= idx; i-- {
	// 	slice[i+1] = slice[i]
	// }
	// slice[idx] = 100
	//--
	// slice = append(slice, 0)
	// copy(slice[idx+1:], slice[idx:])
	// slice[idx] = 100
	//---------------------------------
	fmt.Println("slice ", slice)

}
