package main

import "fmt"

func main() {
	var countryMap = make(map[string]string)
	countryMap["France"] = "巴黎"
	countryMap["Italy"] = "罗马"
	countryMap["Japan"] = "东京"
	countryMap["India"] = "新德里"
	for key, value := range countryMap {
		fmt.Println(key + ":=" + value)
		fmt.Println(countryMap[key])
	}

	delete(countryMap, "Japan")
	fmt.Println("打印删除后的map集合")
	for country := range countryMap {
		fmt.Println(country + "的首都是：" + countryMap[country])
	}

}
