package flyweight

type game struct {
	terrorists        []*player
	counterTerrorists []*player
}

func NewGame() *game {
	return &game{
		terrorists:        make([]*player, 0),
		counterTerrorists: make([]*player, 0),
	}
}

func (c *game) AddTerrorist(dressType string) {
	player := NewPlayer("T", dressType)
	c.terrorists = append(c.terrorists, player)
}

func (c *game) AddCounterTerrorist(dressType string) {
	player := NewPlayer("CT", dressType)
	c.counterTerrorists = append(c.counterTerrorists, player)
}
