package main

import "fmt"

type TZ int

func main() {
	var a TZ //变量a为TZ类型
	//a = 2
	a.alert(100)
	fmt.Println(a)
}

func (t *TZ) alert(i int) {
	*t = TZ(i) //强转
}
