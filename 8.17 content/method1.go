package main

import "fmt"

type A struct {
	s string
}

type B struct {
	b string
}

func main() {
	a := A{}
	a.print()
	b := B{}
	b.print()
}

func (a A) print() {
	fmt.Println("call func_printA test")
}

func (b B) print() {
	fmt.Println("call func_printB test")
}
