package builder

import (
	"builder-pattern-go/pkg/product"
)

type gourmetPizzaBuilder struct {
	crust    string
	toppings []string
	size     int
}

func NewGourmetPizzaBuilder() PizzaBuilder {
	return &gourmetPizzaBuilder{}
}

func (b *gourmetPizzaBuilder) SetCrust(crust string) PizzaBuilder {
	b.crust = crust
	return b
}

func (b *gourmetPizzaBuilder) SetToppings(toppings []string) PizzaBuilder {
	b.toppings = toppings
	return b
}

func (b *gourmetPizzaBuilder) SetSize(size int) PizzaBuilder {
	b.size = size
	return b
}

func (b *gourmetPizzaBuilder) Build() (product.Pizza, error) {
	if err := validatePizza(b.crust, b.size, b.toppings); err != nil {
		return product.Pizza{}, err
	}
	return product.Pizza{
		Crust:    b.crust,
		Toppings: b.toppings,
		Size:     b.size,
	}, nil
}
