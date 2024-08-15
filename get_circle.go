package sprint

type Circle struct {
	Radius    float32
	Diameter  float32
	Area      float32
	Perimeter float32
}

func GetCircle(r float32) Circle {
	const pi float32 = 3.14
	return Circle{r, 2 * r, pi * (r * r), 2 * pi * r}
}
