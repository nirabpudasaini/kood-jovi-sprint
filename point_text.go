package sprint

import "fmt"

type Point struct {
	X    float32
	Y    float32
	Text string
}

func PointText(p Point) Point {
	return Point{p.X, p.Y, fmt.Sprintf("%s at (%f, %f)", p.Text, p.X, p.Y)}
}
