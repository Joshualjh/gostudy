package main

import "fmt"

type Reader interface {
	Read()
}

type Closer interface {
	Close()
}

type File struct {
}

func (f *File) Read() {
	fmt.Println("Read")
}

func (f *File) Close() {
	fmt.Println("Close")
}

func ReadFile(reader Reader) {
	reader.Read()
	if c, err := reader.(Closer); err {
		c.Close()
	}
}
func main() {
	file := &File{}
	ReadFile(file)
}
