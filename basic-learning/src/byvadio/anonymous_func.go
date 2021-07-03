/**
 @author: lin.she
 @date: 2021/6/28
 @note:
**/
package main

import "fmt"

func main() {
	func(a, b int) {
		fmt.Print("这是一个匿名函数")
	}(1, 3)
}
