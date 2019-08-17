package main

import "fmt"

type T struct {
	u_id   int
	u_name string
}

type t T
type tt t

func main() {
	a := t{}
	a.print()
}

func (rt t) print() {
	fmt.Println("call func_print")
}
