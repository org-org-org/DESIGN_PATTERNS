package visitor

type Rectangle struct {
	L int
	B int
}

func (t *Rectangle) Accept(v visitor) {
	v.visitForRectangle(t)
}

func (t *Rectangle) getType() string {
	return "rectangle"
}
