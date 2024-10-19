package packages

const (
	PI = 3.14
)

type Circle struct {
	Radius float64
}

func NewCicle(r float64) *Circle {
	return &Circle{
		Radius: r,
	}
}

func (c *Circle) Cal_Circle_Circum() float64 {
	return (2 * PI * c.Radius)
}

func (c *Circle) Cal_Circle_Area() float64 {
	return (PI * c.Radius * c.Radius)
}
