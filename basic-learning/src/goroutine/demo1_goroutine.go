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

	start := time.Now().UnixNano()

	go printNum()

	for i := 0; i < 100; i++ {
		fmt.Printf("\t主goroutine打印字母：%s\n", "abc")
	}

	end := time.Now().UnixNano()

	fmt.Println("打印结束。。。。。", (end - start))

}

func printNum() {
	for i := 0; i < 100; i++ {
		fmt.Printf("子goroutine打印数字：%d\n", i)
	}
}
