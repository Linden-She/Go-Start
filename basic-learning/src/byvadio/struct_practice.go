/**
 @author: lin.she
 @date: 2021/7/4
 @note:
**/
package main

import "fmt"

func main() {
	stu1 := new(Student)
	stu1.name = "Linden"
	stu1.age = 25
	fmt.Println(stu1)

	// 匿名结构体
	stu2 := struct {
		name string
		age  int
	}{
		name: "linden",
		age:  25,
	}
	fmt.Println(stu2)
}

type Student struct {
	name string
	age  int
}
