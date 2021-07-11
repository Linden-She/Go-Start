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

	ch := make(chan int)
	go sendData(ch)

	for {
		time.Sleep(500 * time.Millisecond)
		v, ok := <-ch
		if !ok {
			fmt.Println("主 goroutine 从通道接收数据结束：", v, ok)
			break
		}
		fmt.Println("主 goroutine 从通道接收数据:", v, ok)
	}

}

func sendData(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}
