/**
 @author: lin.she
 @date: 2021/6/28
 @note:
**/
package main

import "fmt"

func main() {
	recu := getSumRecu(0, 5)
	fmt.Print(recu)
}

func getSumRecu(sum int, a int) int {
	if a <= 0 {
		return sum
	}
	sum += a
	return getSumRecu(sum, a-1)
}
