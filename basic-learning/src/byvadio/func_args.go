/**
 @author: lin.she
 @date: 2021/6/28
 @note:
**/
package main

import "fmt"

func main() {
	getSum(1, 2, 3, 4, 5, 6)
	param := []int{2, 4, 6, 8}
	getSum(param...)
}

func getSum(args ...int) {
	sum := 0
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}
	fmt.Println(sum)
}
