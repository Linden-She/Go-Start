/**
 @author: lin.she
 @date: 2021/7/4
 @note:
**/
package main

import "fmt"

func main() {
	per1 := Person{name: "linden", age: 18, address: Address{city: "xian", state: "changanqu"}}
	fmt.Println(per1)
}

type Person struct {
	name    string
	age     int
	address Address
}

type Address struct {
	city  string
	state string
}
