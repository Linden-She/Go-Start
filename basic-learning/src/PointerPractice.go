package main

import "fmt"

const length = 3

func main() {
	test3()
}

func test3() {
	var a = [length]int{2,6,7}
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
	println( &a)
	var vp *int
	a = 20
	vp = &a
	println(vp)
	println(*vp)
}

