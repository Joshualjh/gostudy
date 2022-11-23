package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3}

	//append를 활용하여 slice 복사
	slice2 := append([]int{}, slice1...)
	// make 와 copy를 활용한 복사
	slice3 := make([]int, len(slice1))
	copy(slice3, slice1)
	//그냥 복사할 경우 슬라이스의 주소값만 복사되어 값을 변경시 초기 슬라이스의 값도 변경된다.
	slice4 := slice1[:]
	// for 문을 이용한 복사 방법도 있음
	slice5 := make([]int, len(slice1), cap(slice1))
	for i, v := range slice1 {
		slice5[i] = v
	}

	slice2[1] = 100
	slice3[1] = 200
	slice5[1] = 3456
	fmt.Println("slice1 :", slice1)
	fmt.Println("slice2 :", slice2)
	fmt.Println("slice3 :", slice3)
	fmt.Println("slice4 :", slice4)
	fmt.Println("slice5 :", slice5)

	slice4[1] = 6000
	fmt.Println("slice1 :", slice1)
	fmt.Println("slice2 :", slice2)
	fmt.Println("slice3 :", slice3)
	fmt.Println("slice4 :", slice4)

}
