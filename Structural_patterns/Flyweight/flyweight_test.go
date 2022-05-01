package t_test

import (
	"DESIGN_PATTERNS/Structural_patterns/Flyweight/flyweight"
	"fmt"
	"testing"
)

func TestFlyweight(t *testing.T) {
	game := flyweight.NewGame()

	//Add Terrorist
	game.AddTerrorist(flyweight.TerroristDressType)
	game.AddTerrorist(flyweight.TerroristDressType)
	game.AddTerrorist(flyweight.TerroristDressType)
	game.AddTerrorist(flyweight.TerroristDressType)

	//Add CounterTerrorist
	game.AddCounterTerrorist(flyweight.CounterTerroristDressType)
	game.AddCounterTerrorist(flyweight.CounterTerroristDressType)
	game.AddCounterTerrorist(flyweight.CounterTerroristDressType)

	dressFactoryInstance := flyweight.GetDressFactorySingleInstance()

	for dressType, dress := range dressFactoryInstance.DressMap {
		fmt.Printf("DressColorType: %s\nDressColor: %s\n", dressType, dress.GetColor())
	}
}
