/**
 @author: lin.she
 @date: 2021/7/3
 @note:
**/
package main

import "fmt"

func main() {
	fmt.Printf("%T\n", add)
	fmt.Printf("%T\n", oper)
	res := oper(10, 20, add)
	fmt.Println(res)

	res2 := oper(10, 20, func(a int, b int) int {
		return a * b
	})
	fmt.Println(res2)

}

func add(a, b int) int {
	return a + b
}

func oper(a int, b int, add func(int, int) int) int {
	fmt.Println(a, b, add)
	return add(a, b)
}
