/**
 @author: lin.she
 @date: 2021/6/23
 @note:
**/
package main

import "fmt"

func main() {

	//s1 := [5]int{1,23,45,6,9}
	//fmt.Println(s1)

	s2 := make([]int, 5, 6)
	s2 = []int{2, 3, 4, 5, 6}
	//fmt.Println(s2)

	s2 = append(s2, 34, 54, 23)
	fmt.Println(s2)

	s3 := make([]int, 0, 0)
	/*for _,value := range s1 {
		s3 = append(s3, value)
	}*/

	fmt.Println(s3)
	copy(s3, s2)
	fmt.Println(s3)

}
