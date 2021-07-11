/**
 @author: lin.she
 @date: 2021/7/11
 @note:
**/
package main

import (
	"fmt"
	"sync"
	"time"
)

var tickets = 10
var wg2 sync.WaitGroup
var mutex sync.Mutex

func main() {
	wg2.Add(4)
	go saleTickets("goroutine 1")
	go saleTickets("goroutine 2")
	go saleTickets("goroutine 3")
	go saleTickets("goroutine 4")

	wg2.Wait()

}

func saleTickets(name string) {
	defer wg2.Done()
	for {
		mutex.Lock()
		if tickets > 0 {
			fmt.Println(name, "窗口售出第", tickets, "张票")
			time.Sleep(1 * time.Second)
			tickets--
		} else {
			mutex.Unlock()
			fmt.Println(name, "窗口售票结束......")
			break
		}
		mutex.Unlock()
	}
}
