package decorator

type CheeseTopping struct {
	Pizza pizza
}

func (c *CheeseTopping) getPrice() int {
	pizzaPrice := c.Pizza.getPrice()
	return pizzaPrice + 10
}
