/**
 @author: lin.she
 @date: 2021/6/24
 @note:
**/
package main

import (
	"fmt"
	"sort"
)

func main() {
	var map1 map[int]string
	var map2 = make(map[string]int)
	map3 := map[string]int{"Go": 98, "Python": 93, "Java": 100}
	fmt.Println(map1)
	fmt.Println(map2)
	fmt.Println(map3)

	map2["Chinese"] = 1998
	val, ok := map2["Chinese"]
	if ok {
		fmt.Println(val, ok)
	}

	if map1 == nil {
		map1 = make(map[int]string)
	}
	map1[1] = "yearstoday"
	map1[2] = "today"
	map1[3] = "tommory"
	map1[4] = "sample"
	map1[5] = "test"
	fmt.Println(map1)
	delete(map1, 2)
	fmt.Println(map1)

	for k, v := range map1 {
		fmt.Println(k, v)
	}
	//顺序打印map中的内容
	// 1.先创建slice保存keys
	slice1 := make([]int, 0, len(map1))
	for k, _ := range map1 {
		slice1 = append(slice1, k)
	}
	sort.Ints(slice1)
	fmt.Println(slice1)
	for _, v := range slice1 {
		fmt.Println(v, map1[v])
	}

}
