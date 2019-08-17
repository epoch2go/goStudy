package main

import "fmt"

/*
声明一个类型的底层类型为int，并实现调用某个方法就递增100
*/

type theT int

func main() {
	var a theT //a=0
	a.Increase(10)
	fmt.Println(a)
}

func (t *theT) Increase(a int) {
	*t = theT(a + 100)
}
