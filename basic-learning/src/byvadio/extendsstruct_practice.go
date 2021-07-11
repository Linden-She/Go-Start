/**
 @author: lin.she
 @date: 2021/7/4
 @note:
**/
package main

import "fmt"

func main() {
	dog1 := Dog{
		Animal: Animal{
			name: "可乐",
			age:  17,
		},
		class: "狗",
	}
	fmt.Println(dog1)
	dog1.eat()
}

type Animal struct {
	name string
	age  int
}

type Dog struct {
	Animal
	class string
}

func (a Animal) eat() {
	fmt.Println(a.name, ":杂食动物，吃anything")
}

func (d Dog) eat() {
	fmt.Println(d.name, ":肉食动物，吃狗粮")
}
