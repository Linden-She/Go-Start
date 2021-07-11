/**
 @author: lin.she
 @date: 2021/7/4
 @note:
**/
package main

import "fmt"

func main() {
	// 数组指针
	arr1 := [4]int{1, 2, 3, 4}
	fmt.Println(arr1)
	var p1 *[4]int
	p1 = &arr1
	fmt.Println(p1)
	fmt.Printf("%p\n", &arr1)
	fmt.Printf("%p\n", p1)
	p1[0] = 100
	(*p1)[1] = 200
	fmt.Println(arr1)

	// 指针数组
	a := 1
	b := 2
	c := 3
	d := 4
	fmt.Println(a, b, c, d)
	arr2 := [4]int{a, b, c, d}
	fmt.Println(arr2)
	arr3 := [4]*int{&a, &b, &c, &d}
	fmt.Println(arr3)

	*arr3[0] = 100
	fmt.Println(a)

}
