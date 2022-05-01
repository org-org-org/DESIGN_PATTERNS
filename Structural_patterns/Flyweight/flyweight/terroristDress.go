package flyweight

type terroristDress struct {
	color string
}

func (t *terroristDress) GetColor() string {
	return t.color
}

func NewTerroristDress() *terroristDress {
	return &terroristDress{color: "red"}
}

type counterTerroristDress struct {
	color string
}

func (c *counterTerroristDress) GetColor() string {
	return c.color
}

func NewCounterTerroristDress() *counterTerroristDress {
	return &counterTerroristDress{color: "green"}
}
