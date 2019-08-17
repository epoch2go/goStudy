package main

import "fmt"

/*
请实现一个函数，将一个字符串中的每个空格替换成“%20”。例如，当字符串为We Are Happy.则经过替换之后的字符串为We%20Are%20Happy。
*/

func main() {
	s := "We Are Happy"
	sl := make([]string, len(s))
	//fmt.Println(sl)
	for i, v := range s {
		if v != 32 {
			sl[i] = string(v)
		}
		//fmt.Print(v, "\t")
		if v == 32 {
			sl[i] = "%20"
		}
	}
	for k, _ := range sl {
		fmt.Print(sl[k])
	}
}
