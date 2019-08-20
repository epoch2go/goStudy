package main

import "fmt"

func main() {
	var temp int
	s:=[]int{11,22,4,551,2,66,123,444}
	for i:=0;i< len(s)-1;i++  {
		for j:= len(s)-1;j>i;j--  {
			if s[j] < s[j-1] {
				temp=s[j]
				s[j]=s[j-1]
				s[j-1]=temp
			}
		}
		fmt.Println("第",i+1,"次排序：",s)
	}
	fmt.Println(s)
}
