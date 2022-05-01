package abstractFactory

type (
	nike struct {
	}
	nikeShoe struct {
		shoe
	}
	nikeShirt struct {
		shirt
	}
)

func (n *nike) MakeShoe() IShoe {
	return &nikeShoe{
		shoe: shoe{
			logo: "nike",
			size: 14,
		},
	}
}

func (n *nike) MakeShirt() IShirt {
	return &nikeShirt{
		shirt: shirt{
			logo: "nike",
			size: 14,
		},
	}
}
