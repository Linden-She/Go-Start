package main

import (
	"fmt"
	"strconv"
)

const (
	name = "Linden"
	age = 25
	hobby = "basketball"
)

func main() {
	var boo bool = true
	fmt.Print(boo)
	var num1 = 5
	fmt.Print(num1)
	var num2 float32 = 5.234
	fmt.Println(num2)
	var arr = [3]int{1,3,5}
	fmt.Println(arr)
	const LENGTH = 5
	fmt.Println(LENGTH)

	fmt.Println("name:" + name + "  age:" + strconv.Itoa(age) + "  hobby:" + hobby )
}
