/**
 @author: lin.she
 @date: 2021/7/11
 @note:
**/
package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {

	go fun1()
	go fun2()
	wg.Add(2)

	wg.Wait()
	fmt.Println("main线程执行完毕")

}

func fun1() {
	for i := 0; i < 10; i++ {
		fmt.Println("子线程fun1打印数字：", i)
	}
	wg.Done()
}

func fun2() {
	for i := 0; i < 10; i++ {
		fmt.Println("子线程fun2打印数字：", i)
	}
	wg.Done()
}
