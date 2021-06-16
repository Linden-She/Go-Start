package example

import "testing"

func Perimeter(a float64, b float64) float64 {
	return 2 * (a + b)
}

func Area(a float64, b float64) float64 {
	return a * b
}

func TestPerimeter(t *testing.T) {
	got := Perimeter(10.0, 10.0)
	want := 40.0
	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	got := Area(12.0, 6.0)
	want := 72.0
	if got != want {
		t.Errorf("got %f want %f", got, want)
	}
}

/**
 * @Author lin.she
 * @Description 计算面积结构体
 * @Date 9:48 2021/6/15
 * @Param
 * @return
 **/
type Reactangle struct {
	Width  float64
	Height float64
}

func PerimeterReact(reactangle Reactangle) float64 {
	return 2 * (reactangle.Width + reactangle.Height)
}

func AreaReact(reactangle Reactangle) float64 {
	return reactangle.Width * reactangle.Height
}

func TestReact(t *testing.T) {
	reactangle := Reactangle{10.0, 10.0}
	got := PerimeterReact(reactangle)
	want := 40.0
	if got != want {
		t.Errorf("got %f want %f", got, want)
	}
}

func TestAreaReact(t *testing.T) {
	reactangle := Reactangle{12.0, 6.0}
	got := AreaReact(reactangle)
	want := 64.0
	if got != want {
		t.Errorf("got %f want %f", got, want)
	}
}
