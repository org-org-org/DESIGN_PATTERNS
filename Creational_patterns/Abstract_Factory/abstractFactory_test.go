package t_test

import (
	"DESIGN_PATTERNS/Creational_patterns/Abstract_Factory/abstractFactory"
	"fmt"
	"testing"
)

func printShoeDetails(s abstractFactory.IShoe) {
	fmt.Printf("Logo: %s\n", s.GetLogo())
	fmt.Printf("Size: %d\n", s.GetSize())
}

func printShirtDetails(s abstractFactory.IShirt) {
	fmt.Printf("Logo: %s\n", s.GetLogo())
	fmt.Printf("Size: %d\n", s.GetSize())
}

func TestAbstractFactory(t *testing.T) {
	adidasFactory, _ := abstractFactory.GetSportsFactory("adidas")
	nikeFactory, _ := abstractFactory.GetSportsFactory("nike")

	nikeShoe := nikeFactory.MakeShoe()
	nikeShirt := nikeFactory.MakeShirt()

	adidasShoe := adidasFactory.MakeShoe()
	adidasShirt := adidasFactory.MakeShirt()

	printShoeDetails(nikeShoe)
	printShirtDetails(nikeShirt)

	printShoeDetails(adidasShoe)
	printShirtDetails(adidasShirt)
}
