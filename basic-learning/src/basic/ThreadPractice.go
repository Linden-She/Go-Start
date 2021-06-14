package main

import (
	"fmt"
	"time"
)

func main() {
	/*go say("world")
	say("hellow")*/
	/*a := []int{2,5,3,7,5,9,12,1}
	c := make(chan int)
	go sumCaculate(a[:len(a)/2], c)
	go sumCaculate(a[len(a)/2:], c)
	x,y := <-c, <-c
	fmt.Println(x, y, x + y)*/
	c := make(chan int, 10)
	go fabonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(s, i)
	}
}

func sumCaculate(num []int, c chan int) {
	sum := 0
	for _, value := range num {
		sum += value
	}
	c <- sum
}

func fabonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}
