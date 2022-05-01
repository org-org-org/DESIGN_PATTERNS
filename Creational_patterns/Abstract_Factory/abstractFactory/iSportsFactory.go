package abstractFactory

import "fmt"

type ISportsFactory interface {
	MakeShoe() IShoe
	MakeShirt() IShirt
}

func GetSportsFactory(brand string) (ISportsFactory, error) {
	if brand == "adidas" {
		return &adidas{}, nil
	}

	if brand == "nike" {
		return &nike{}, nil
	}

	return nil, fmt.Errorf("wrong brand type passed")
}
