package main

import "fmt"

func main() {
	var temp int
	s:=[]int{11,22,4,551,2,66,123,444}

	for i := 1; i< len(s);i++  {
		//temp=s[i]
		a:=i
		for j:=i-1;j>=0;j--  {
			if s[a]<s[j] {
				temp=s[a]
				s[a]=s[j]
				s[j]=temp
				a--
			}
		}
		fmt.Println("第",i,"次排序：",s)
	}
}