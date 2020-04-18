package main

import (
	"fmt"
	"math"
)

type Point struct{ X, Y float64 }

func main() {
	p := Point{1, 2}
	q := Point{4, 6}
	fmt.Println(Distance(p, q))
	fmt.Println(p.Diatance(q))
}

func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (p Point) Diatance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}
