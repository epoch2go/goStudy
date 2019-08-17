package main

import "fmt"

type user struct {
	u_id   int
	u_name string
}

func main() {
	a := user{
		u_id:   123,
		u_name: "name1",
	}
	fmt.Println(a)
	fmt.Printf("%p\n", &a)
	a.alterName()
	fmt.Println(a)
	fmt.Printf("%p", &a)
}

func (a *user) alterName() {
	a.u_name = "altername"
	a.u_id = 321
	fmt.Println("call func_alterName")
}
