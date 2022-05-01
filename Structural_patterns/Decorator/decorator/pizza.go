package decorator

type pizza interface {
	getPrice() int
}

type VeggeMania struct {
}

func (p *VeggeMania) getPrice() int {
	return 15
}
