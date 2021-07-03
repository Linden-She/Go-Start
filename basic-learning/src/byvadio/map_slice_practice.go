/**
 @author: lin.she
 @date: 2021/6/25
 @note:
**/
package main

import "fmt"

func main() {
	map1 := make(map[string]string)
	map1["name"] = "Java"
	map1["age"] = "20"
	map1["sex"] = "male"
	map1["address"] = "BeiJing"
	map2 := map[string]string{"name": "Ruby", "age": "15", "sex": "female", "address": "ShenZhen"}
	map3 := map[string]string{"name": "Go", "age": "18", "sex": "male", "address": "Xian"}

	fmt.Println(map1)
	fmt.Println(map2)
	fmt.Println(map3)

	s1 := make([]map[string]string, 0, 3)
	s1 = append(s1, map1)
	s1 = append(s1, map2)
	s1 = append(s1, map3)

	for i, val := range s1 {
		fmt.Printf("第 %d 个元素的信息是：\n", i+1)
		fmt.Printf("\t姓名：%s\n", val["name"])
		fmt.Printf("\t年龄：%s\n", val["age"])
		fmt.Printf("\t性别：%s\n", val["sex"])
		fmt.Printf("\t地址：%s\n", val["address"])
	}

	map4 := make(map[int]string)
	fmt.Println(map4)

}
