/**
 @author: lin.she
 @date: 2021/7/11
 @note:
**/
package main

import (
	"fmt"
	"strconv"
)

func main() {

	ch1 := make(chan string, 2)

	go sendData1(ch1)

	for {
		val, ok := <-ch1
		if !ok {
			fmt.Println("通道内容读取完毕", val, ok)
			break
		}
		fmt.Printf("\t主 goroutine 接收数据 %s\n", val)
	}

	fmt.Println("main excute over .......")

}

func sendData1(ch1 chan string) {
	for i := 0; i < 10; i++ {
		ch1 <- "数据：" + strconv.Itoa(i)
		fmt.Printf("子 goroutine 写入第 %d 个数据\n", i)
	}
	close(ch1)
}
