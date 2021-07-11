/**
 @author: lin.she
 @date: 2021/7/4
 @note:
**/
package main

import "fmt"

func main() {
	m1 := Mouse{name: "Logit"}
	m1.start()
	m1.end()
	f1 := FlashDisk{name: "what"}
	textInterface(f1)
}

type USB interface {
	start()
	end()
}

type Mouse struct {
	name string
}

type FlashDisk struct {
	name string
}

func (m Mouse) start() {
	fmt.Println("鼠标开始接入，click......")
}

func (m Mouse) end() {
	fmt.Println("鼠标结束接入，esc......")
}

func (f FlashDisk) start() {
	fmt.Println(f.name, "接入，开始传输数据......")
}

func (f FlashDisk) end() {
	fmt.Println(f.name, "结束数据传输......")
}

func textInterface(u USB) {
	u.start()
	u.end()
}
