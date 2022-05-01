package visitor

type Circle struct {
	Radius int
}

func (c *Circle) Accept(v visitor) {
	v.visitForCircle(c)
}

func (c *Circle) getType() string {
	return "Circle"
}
