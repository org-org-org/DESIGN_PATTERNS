package t_test

import (
	"DESIGN_PATTERNS/Creational_patterns/Factor_Method/Factor"
	"fmt"
	"testing"
)

func printDetails(g Factor.IGun) {
	fmt.Printf("Gun: %s\n", g.GetName())
	fmt.Printf("Power: %d\n", g.GetPower())
}

func TestFactor(t *testing.T) {
	ak47, _ := Factor.GetGun("ak47")
	musket, _ := Factor.GetGun("musket")

	printDetails(ak47)
	printDetails(musket)
}
