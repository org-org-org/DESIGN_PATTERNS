package t_test

import (
	"DESIGN_PATTERNS/Structural_patterns/Decorator/decorator"
	"fmt"
	"testing"
)

func TestDecorator(t *testing.T) {
	pizza := &decorator.VeggeMania{}
	// Add cheese topping
	pizzaWithCheese := &decorator.CheeseTopping{
		Pizza: pizza,
	}

	// Add tomato topping
	pizzaWithCheeseAndTomato := &decorator.TomatoTopping{
		Pizza: pizzaWithCheese,
	}

	fmt.Printf("Price of veggeMania with tomato and cheese topping is %d\n", pizzaWithCheeseAndTomato.GetPrice())
}
