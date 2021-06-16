/**
 @author: lin.she
 @date: 2021/6/16
 @note:
**/
package example

import (
	"fmt"
	"math"
	"testing"
)

type RectangleNew struct {
	Width  float64
	Height float64
}

func (r RectangleNew) AreaMethod() float64 {
	return r.Width * r.Height
}

type CircleNew struct {
	Radius float64
}

func (c CircleNew) AreaMethod() float64 {
	return c.Radius * c.Radius * math.Pi
}

func TestMethod(t *testing.T) {
	reactangle := RectangleNew{10.0, 10.0}
	area := reactangle.AreaMethod()
	fmt.Println(area)
}
