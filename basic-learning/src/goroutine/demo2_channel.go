/**
 @author: lin.she
 @date: 2021/7/11
 @note:
**/
package main

import (
	"fmt"
)

func main() {

	ch1 := make(chan bool)

	go func() {

		for i := 0; i < 10; i++ {
			fmt.Println("子 goroutine 执行", i)
		}
		data := <-ch1
		fmt.Println("子 goroutine 执行完毕。。。。。。", data)
	}()
	ch1 <- true

	fmt.Println("主 goroutine 执行完毕......")

}
