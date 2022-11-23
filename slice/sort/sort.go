package main

import (
	"fmt"
	"sort"
)

type Student struct {
	Name string
	Age  int
}

type Students []Student

func (s Students) Len() int           { return len(s) }
func (s Students) Less(i, j int) bool { return s[i].Age < s[j].Age }
func (s Students) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func main() {
	// 구조체 정렬방법
	s := []Student{
		{"화랑", 31},
		{"백두산", 52},
		{"류", 42},
		{"켄", 38},
		{"송하나", 18},
	}
	sort.Sort(Students(s))
	fmt.Println(s)
	//----------------------
	// 기본 정렬 방법
	// slice := []int{3, 6, 8, 4, 7, 1, 2, 5}
	// sort.Ints(slice)
	// fmt.Println(slice)

}
