package main

import "fmt"

type IReadWriter interface {
	Read(b *byte, cb int) int
	Write(b *byte, cb int) int
}

type Class struct {
	a int
}

func (class Class) Read(b *byte, cb int) int {
	fmt.Println("Class.Read: ", class.a)
	return cb
}

func (class Class) Write(b *byte, cb int) int {
	fmt.Println("Class.Write: ", class.a)
	return cb
}

func main() {
	fmt.Println("hello world")
	var p IReadWriter
	p = Class{10}
	p.Read(nil, 5)
}
