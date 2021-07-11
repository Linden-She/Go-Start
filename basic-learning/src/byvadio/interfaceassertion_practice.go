/**
 @author: lin.she
 @date: 2021/7/5
 @note:
**/
package main

import (
	"fmt"
	"math"
)

func main() {
	t := Triangle{3, 4, 5}
	fmt.Println(t.peri())
	fmt.Println(t.area())

	c := Circle{5}
	fmt.Println(c.peri())
	fmt.Println(c.area())

	caculateShap(t)
	caculateShap(c)

	getType(t)
	getType(c)

	cp := &c
	getType(cp)

	getTypeBySwitch(t)
}

func getTypeBySwitch(s Shap) {
	switch ins := s.(type) {
	case Triangle:
		fmt.Println("三角形，边长分别是：", ins.a, ins.b, ins.c)
	case Circle:
		fmt.Println("圆形，半径为：", ins.r)
	}
}

func getType(s Shap) {
	if ins, ok := s.(Triangle); ok {
		fmt.Println(ok)
		fmt.Println("是三角形，边长分别为：", ins.a, ins.b, ins.c)
	} else if ins, ok := s.(Circle); ok {
		fmt.Println("是圆形，半径为：", ins.r)
	} else if ins, ok := s.(*Circle); ok {
		fmt.Printf("ins-----------%T,%p\n", ins, &ins)
		fmt.Printf("s-----------%T,%p\n", s, &s)
	} else {
		fmt.Println("未知类型。。。。")
	}
}

func caculateShap(s Shap) {
	fmt.Println("图形的周长：%2f,面积：%2f", s.peri(), s.area())
}

type Shap interface {
	peri() float64
	area() float64
}

type Triangle struct {
	a, b, c float64
}

func (t Triangle) peri() float64 {
	return t.a + t.b + t.c
}

func (t Triangle) area() float64 {
	p := t.peri()
	return math.Sqrt(p * (p - t.a) * (p - t.b) * (p - t.c))
}

type Circle struct {
	r float64
}

func (c Circle) peri() float64 {
	return 2 * c.r * math.Pi
}

func (c Circle) area() float64 {
	return math.Pi * c.r * c.r
}
