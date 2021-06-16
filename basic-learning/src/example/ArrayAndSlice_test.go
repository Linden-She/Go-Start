package example

import (
	"fmt"
	"testing"
)

func Sum(numbers [5]int) int {
	sum := 0
	for i := 0; i < len(numbers); i++ {
		sum += numbers[i]
	}
	return sum
}

func TestSum(t *testing.T) {
	numbers := [5]int{1, 2, 3, 4, 5}
	got := Sum(numbers)
	want := 15
	if got != want {
		t.Errorf("got %d want %d given, %v", got, want, numbers)
	}
}

func SumSlice(numbers []int) int {
	sum := 0
	for i := 0; i < len(numbers); i++ {
		sum += numbers[i]
	}
	return sum
}

func TestSumBySlce(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	slice := SumSlice(numbers)
	fmt.Println(slice)
}
