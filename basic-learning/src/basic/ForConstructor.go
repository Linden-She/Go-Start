package main

func main() {
	var a = [6]int{1, 4, 2, 3, 6}
	var sum = 0
	for i := 0; i < 6; i++ {
		sum += a[i]
	}
	println(sum)
}
