package main

import "testing"

func TestPerimeter(t *testing.T) {
	got := Perimeter(Rectangle{Width: 10.0, Height: 10.0})
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	// checkArea := func(t testing.TB, shape Shape, want float64) {
	// 	t.Helper()
	// 	got := shape.Area()
	// 	if got != want {
	// 		t.Errorf("got %g want %g", got, want)
	// 	}
	// }

	// t.Run("rectangles", func(t *testing.T) {
	// 	rectangle := Rectangle{Width: 12.0, Height: 6.0}
	// 	want := 72.0
	// 	checkArea(t, rectangle, want)
	// })
	// t.Run("circles", func(t *testing.T) {
	// 	circle := Circle{10}
	// 	want := 314.1592653589793
	// 	checkArea(t, circle, want)
	// })

}
