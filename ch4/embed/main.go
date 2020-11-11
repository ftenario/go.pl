package main

import "fmt"

type Point struct {
	x, y int
}

type Circle struct {
	Point
	Radius int
}

type Wheel struct {
	Circle
	Spokes int
}

func main() {
	var w Wheel
	w = Wheel{
		Circle: Circle{
			Point:  Point{x: 8, y: 10},
			Radius: 5,
		},
		Spokes: 20,
	}

	fmt.Printf("%+v\n", w)
	w.x = 100
	w.Circle.Radius = 20
	fmt.Printf("%+v\n", w)

}
