package decorator

type TomatoTopping struct {
	Pizza pizza
}

func (c *TomatoTopping) GetPrice() int {
	pizzaPrice := c.Pizza.getPrice()
	return pizzaPrice + 7
}
