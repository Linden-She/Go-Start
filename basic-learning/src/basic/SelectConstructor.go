package main

func main() {
	var a, b, c chan int
	select {
	case a <- 90:
		println(a)
	case b <- 80:
		println(b)
	case c <- 60:
		println(c)
	default:
		println(1)
	}
}
