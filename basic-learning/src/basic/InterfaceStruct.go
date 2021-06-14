package main

import "fmt"

type phone interface {
	call()
}

type NokiaPhone struct {
}

func (nokiaPhone NokiaPhone) call() {
	fmt.Println("I am Nokia phone")
}

type Iphone struct {
}

func (iphone Iphone) call() {
	fmt.Println("I am Iphone")
}

func main() {
	nokia := new(NokiaPhone)
	nokia.call()

	iPhone := new(Iphone)
	iPhone.call()

}
