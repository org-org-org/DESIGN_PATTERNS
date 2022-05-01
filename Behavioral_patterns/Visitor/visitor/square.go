package visitor

type Square struct {
	Side int
}

func (s *Square) Accept(v visitor) {
	v.visitForSquare(s)
}

func (s *Square) getType() string {
	return "Square"
}
