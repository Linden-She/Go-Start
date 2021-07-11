/**
 @author: lin.she
 @date: 2021/7/4
 @note:
**/
package main

import "fmt"

func main() {
	addr, res := increment()
	fmt.Println("函数外部变量i地址", addr)
	fmt.Printf("%T\n", res)
	fmt.Println(res)
	i := res()
	fmt.Println("函数外部变量i：", &i)

	_, res2 := increment()
	i2 := res2()
	fmt.Println(i2)
}

func increment() (*int, func() int) {
	i := 1
	addr := &i
	fmt.Printf("%T\n", addr)
	fmt.Println("函数内部变量i：", &i)
	res := func() int {
		i++
		fmt.Println("函数内内部变量i：", &i)
		return i
	}
	return addr, res
}
