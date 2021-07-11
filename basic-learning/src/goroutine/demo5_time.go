/**
 @author: lin.she
 @date: 2021/7/11
 @note:
**/
package main

import (
	"fmt"
	"time"
)

func main() {

	//timer := time.NewTimer(3 * time.Second)
	//
	//fmt.Printf("%T\n", timer)
	//fmt.Println(time.Now())
	//
	//ch3 := timer.C
	//fmt.Println("1----------------------")
	//fmt.Println("主 goroutine 执行结束。。。。。", <-ch3)

	timer2 := time.NewTimer(2 * time.Second)

	go func() {
		<-timer2.C
		fmt.Println("子 goroutine 执行完毕")
	}()

	time.Sleep(5 * time.Second)

	fmt.Println("主 goroutine 执行完毕")

}
