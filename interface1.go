package main

import "fmt"

type Phone interface {
	call()
}

type Iphone8 struct {
}

type IphoneX struct {
}

func main() {
	var p Phone
	p = Iphone8{}
	p.call()
	p = IphoneX{}
	p.call()
}

func (i Iphone8) call() {
	fmt.Println("iphone8 call")
}
func (i IphoneX) call() {
	fmt.Println("iphoneX call")
}
