package main

import "fmt"

func main() {

	nums := []int{2, 3, 5, 6, 78, 9}
	sum := 0
	for index, value := range nums {
		fmt.Println("index:", index)
		fmt.Println("value", value)
		sum += value
	}
	fmt.Println("sum:", sum)

}
