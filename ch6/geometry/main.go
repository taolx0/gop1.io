package geometry

import (
	"math"
)

type Point struct{ X, Y float64 }

type Path []Point

/*func main() {
	p := Point{1, 2}
	q := Point{4, 6}
	fmt.Println(Distance(p, q))
	fmt.Println(p.Diatance(q))

	perim := Path{
		{1, 1},
		{5, 1},
		{5, 4},
		{1, 1},
	}
	fmt.Println(perim.Distance())
	/*one*/
/*		r := &Point{1, 2}
		r.ScaleBy(2)
		fmt.Println(*r)*/

	/*two*/
/*		r := Point{1, 2}
		pptr := &r
		pptr.ScaleBy(2)
		fmt.Println(r)*/

	/*three*/
/*	r := Point{1, 2}
	(&r).ScaleBy(2)
	fmt.Println(r)*/
/*}*/

func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Diatance(path[i])
		}
	}
	return sum
}

func (p Point) Diatance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}
