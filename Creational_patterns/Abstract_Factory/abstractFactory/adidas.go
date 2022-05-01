package abstractFactory

type (
	adidas struct {
	}
	adidasShirt struct {
		shirt
	}
	adidasShoe struct {
		shoe
	}
)

func (a *adidas) MakeShoe() IShoe {
	return &adidasShoe{
		shoe: shoe{
			logo: "adidas",
			size: 14,
		},
	}
}

func (a *adidas) MakeShirt() IShirt {
	return &adidasShirt{
		shirt: shirt{
			logo: "adidas",
			size: 14,
		},
	}
}
