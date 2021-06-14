package main

func main() {
	var arr = []int{3, 4, 6, 21, 87, 34}
	max := findMax(arr)
	println(max)
}

func findMax(arr []int) int {
	var res = 0
	for i := 0; i < len(arr); i++ {
		if arr[i] > res {
			res = arr[i]
		}
	}
	return res
}
