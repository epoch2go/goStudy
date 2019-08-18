package main

import (
	"fmt"
	"strconv"
)

/*
大家都知道斐波那契数列，现在要求输入一个整数n，请你输出斐波那契数列的第n项（从0开始，第0项为0）。
*/

func main() {
	var s string
	fmt.Println("enter:")
	fmt.Scanf("%s", &s)
	a, _ := strconv.Atoi(s)
	if a < 0 || a > 39 {
		fmt.Println("输入有误")
		return
	}
	Fibonacci(a)
}

func Fibonacci(i int) {
	s := make([]int, 40)
	s[0] = 0
	s[1] = 1
	for j := 2; j < 40; j++ {
		s[j] = s[j-1] + s[j-2]
	}
	fmt.Print(s[i])
}
