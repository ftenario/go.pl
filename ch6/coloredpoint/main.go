package main

import (
	"fmt"
	"image/color"
	"math"
)

type Point struct {
	X, Y float64
}

type ColoredPoint struct {
	Point
	Color color.RGBA
}

func New() *ColoredPoint {
	return &ColoredPoint{}
}

func (p *Point) Distance(n Point) float64 {
	total := ((n.X - p.X) * (n.X - p.X)) + ((n.Y - p.Y) * (n.Y - p.Y))
	return math.Sqrt(total)
}

func (p *Point) ScaleBy(n int) {
	p.X = p.X * float64(n)
	p.Y = p.Y * float64(n)
}

func main() {
	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}
	var p = ColoredPoint{
		Point{1, 1},
		red,
	}
	var q = ColoredPoint{
		Point{5, 4},
		blue,
	}
	fmt.Println(p.Distance(q.Point))
	p.ScaleBy(2)
	q.ScaleBy(2)
	fmt.Println(p.Distance(q.Point))

}
