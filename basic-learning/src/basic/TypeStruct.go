package main

import "fmt"

type Book struct {
	title  string
	author string
	price  float32
}

func main() {
	var sh Book
	sh = Book{title: "shuihuzhuan", author: "shinaian", price: 2.53}
	fmt.Println(sh)
}
