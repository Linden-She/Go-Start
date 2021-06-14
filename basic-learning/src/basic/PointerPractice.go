package main

import "fmt"

const length = 3

func main() {
	var a = 10
	var b = 20
	var x = &a
	var y = &b
	swap(x, y)
	fmt.Println("x的值是： %x\n", *x)
	fmt.Println("y的值是：%x\n", *y)
	fmt.Println("x的地址是：%x\n", x)
	fmt.Println("y的地址是：%x\n", y)
	fmt.Println("a的地址是：%x\n", &a)
	fmt.Println("b的地址是：%x\n", &b)
}

func swap2(x int, y int) {
	fmt.Println(&x)
	fmt.Println(&y)
	var tmp int
	tmp = x
	x = y
	y = tmp
}

func swap(x *int, y *int) {
	println(x)
	println(y)
	var tmp int
	tmp = *x
	//println(tmp)
	*x = *y
	//println(x)
	*y = tmp
	//println(y)
}

func test4() {
	var a int
	var b *int
	var c **int
	var d ***int
	b = &a
	println(b)
	c = &b
	println(c)
	d = &c
	println(d)
}

func test3() {
	var a = [length]int{2, 6, 7}
	var ptr [length]*int
	for i := 0; i < length; i++ {
		ptr[i] = &a[i]
	}
	for i := range ptr {
		println(*ptr[i])
	}
}

func test2() {
	var pointer *int
	fmt.Printf(": %x\n", pointer)
}

func test1() {
	var a = 10
	println(a)
	println(&a)
	var vp *int
	a = 20
	vp = &a
	println(vp)
	println(*vp)
}
