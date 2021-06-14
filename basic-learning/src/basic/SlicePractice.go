package main

import "fmt"

func main() {
	//var a = make([]string, 4, 8)
	//fmt.Println(a)
	//fmt.Println(len(a))
	//fmt.Println(cap(a))

	numbers := []int{2, 4, 6, 5, 3}
	/*printSlice(numbers)
	//打印索引1(包含)~4(不包含)
	fmt.Println("numbers[1:4] == ", numbers[1:4])
	//打印索引3以前的数据
	fmt.Println("index小于3 == ",numbers[:3])
	fmt.Println("index大于3 == ",numbers[3:])*/

	//追加切片
	numbers = append(numbers, 0)
	fmt.Println(numbers)

	numbers = append(numbers, 1)
	fmt.Println(numbers)

	numbers = append(numbers, 1, 2, 3, 4)
	fmt.Println(numbers)
}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
