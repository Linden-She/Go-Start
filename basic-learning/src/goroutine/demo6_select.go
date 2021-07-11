/**
 @author: lin.she
 @date: 2021/7/11
 @note:
**/
package main

import "fmt"

func main() {

	ch1 := make(chan bool)
	ch2 := make(chan bool)

	go func() {
		ch1 <- false
	}()

	select {
	case v1 := <-ch1:
		fmt.Println("ch1 接收到通信", v1)
	case v2, ok := <-ch2:
		if ok {
			fmt.Println("ch2 接收到通信", v2, ok)
		} else {
			fmt.Println("ch2 已经关闭了。。。。。。")
		}
	}

}
