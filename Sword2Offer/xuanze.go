package main

import "fmt"

func main() {
	var min int
	s:=[]int{11,22,4,551,2,66,123,444}
	for i:=0;i< len(s)-1; i++ {
		for j:=i+1;j<len(s);j++{
			if s[i]>s[j] {
				min=s[j]
				s[j]=s[i]
				s[i]=min
			}
		}
		fmt.Println("第",i+1,"次排序：",s)
	}
	fmt.Println(s)
}