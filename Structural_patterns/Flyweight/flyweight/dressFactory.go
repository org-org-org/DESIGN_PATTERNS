package flyweight

import "fmt"

const (
	// TerroristDressType terrorist dress type
	TerroristDressType = "tDress"
	// CounterTerroristDressType CT dress type
	CounterTerroristDressType = "ctDress"
)

type dress interface {
	GetColor() string
}

var (
	dressFactorySingleInstance = &dressFactory{
		DressMap: make(map[string]dress),
	}
)

type dressFactory struct {
	DressMap map[string]dress
}

func (d *dressFactory) GetDressByType(dressType string) (dress, error) {
	if d.DressMap[dressType] != nil {
		return d.DressMap[dressType], nil
	}
	if dressType == TerroristDressType {
		d.DressMap[dressType] = NewTerroristDress()
		return d.DressMap[dressType], nil
	}
	if dressType == CounterTerroristDressType {
		d.DressMap[dressType] = NewCounterTerroristDress()
		return d.DressMap[dressType], nil
	}

	return nil, fmt.Errorf("wrong dress type passed")
}

func GetDressFactorySingleInstance() *dressFactory {
	return dressFactorySingleInstance
}
