package main

func main() {
	var grade = 78
	var res = "A"
	switch grade {
	case 90:
		res = "A"
	case 80:
		res = "B"
	case 70:
		res = "C"
	}
	println(res)
}
