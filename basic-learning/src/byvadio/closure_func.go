/**
 @author: lin.she
 @date: 2021/7/4
 @note:
**/
package main

import "fmt"

func main() {
	res := increment()
	fmt.Errorf("%T\n", res)
	fmt.Println(res)
	i := res()
	fmt.Println("函数外部变量i：", &i)

	res2 := increment()
	i2 := res2()
	fmt.Println(i2)
}

func increment() func() int {
	i := 1
	fmt.Println("函数内部变量i：", &i)
	res := func() int {
		i++
		fmt.Println("函数内内部变量i：", &i)
		return i
	}
	return res
}
