package main

import "fmt"

/*
type name struct{}
结构体是一种值类型
结构名不同 类型不同
*/

//type test1 struct {
//}
//
//type test2 struct {
//	i int
//}
//
//type test3 struct {
//	a int
//	b string
//}
//
//type test4 struct {
//	c int
//	d string
//}
//
//type test5 struct {
//	e int
//	f string
//}
//
//type student struct {
//	id   int
//	name string
//	age  int
//}
//
//func A(t4 test4) {
//	t4.c = 3
//	t4.d = "A"
//}
//
//func B(t5 *test4) {
//	t5.c = 3
//	t5.d = "A"
//}
//
//func C(t5 *test5) {
//	t5.e = 333
//	t5.f = "AAA"
//}

func main() {
	//s1 := test1{}
	//fmt.Println(s1) //{}
	//
	//s2 := test2{}
	//fmt.Println(s2) //{0}
	//s2.i = 1
	//fmt.Println(s2) //{1}
	//
	//s3 := test3{a: 1, b: "test3"}
	//s4 := test3{a: 2, b: "test3s4"}
	//fmt.Println(s3) //{1 test3}
	//fmt.Println(s4) //{2 test3s4}
	//
	//fcA := test4{c: 1, d: "q"}
	//A(fcA)
	//fmt.Println(fcA) //{1 q}
	//
	//fcB := test4{c: 1, d: "q"}
	//B(&fcB)
	//fmt.Println(fcB) //{3 A}
	//fcB.c = 2
	//fmt.Println(fcB) //{2 A}
	//
	//fcC := &test5{e: 6, f: "l"}
	//C(fcC)
	//fmt.Println(fcC)
	//fcC.e = 00
	//fmt.Println(fcC)
	//
	//stu1 := student{
	//	id:   1,
	//	name: "stu1",
	//	age:  11,
	//}
	//fmt.Println(stu1)
	//
	//stu2 := student{
	//	id:   2,
	//	name: "stu2",
	//	age:  11,
	//}
	//fmt.Println(stu2)
	//
	//stu3 := &student{
	//	id:   3,
	//	name: "stu3",
	//	age:  12,
	//}
	//fmt.Println(stu3)
	//
	//stu3.id = 4
	//stu2.id = 1
	//fmt.Println(stu1, stu2, stu3)

	//a := struct {
	//	a int
	//	b string
	//}{
	//	a: 1,
	//	b: "q",
	//}

	//匿名结构
	//type t struct {
	//	a, b int
	//	v    struct {
	//		c, d int
	//	}
	//}
	//a := t{
	//	a: 1,
	//	b: 2,
	//}
	//fmt.Println(a)
	//a.v.c = 3 //对下一层结构体的调用
	//a.v.d = 4
	//fmt.Println(a)

	//type user struct {
	//	u_id       int
	//	u_name     string
	//	u_password string
	//	sex        int // male=0 female=1
	//	if_limit   string
	//}
	//
	//type vip_user struct {
	//	user
	//	v_level string
	//}
	//
	//type limit_user struct {
	//	user
	//	limit_time string
	//}
	//
	//u1 := vip_user{
	//	user: user{
	//		u_id: 1, u_name: "user1", u_password: "123", sex: 1,
	//	},
	//	v_level: "v1",
	//}
	//fmt.Println(u1.user.u_name)

	type A struct {
		name string
	}

	type B struct {
		A
		name string
	}

	a := B{
		A:    A{name: "Aname"},
		name: "Bname",
	}

	fmt.Println(a.name)
	fmt.Println(a.A.name)

}
