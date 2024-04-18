package builder

import (
	"builder-pattern-go/pkg/product"
)

type pizzaBuilder struct {
	crust    string
	toppings []string
	size     int
}

func NewPizzaBuilder() PizzaBuilder {
	return &pizzaBuilder{}
}

func (b *pizzaBuilder) SetCrust(crust string) PizzaBuilder {
	b.crust = crust
	return b
}

func (b *pizzaBuilder) SetToppings(toppings []string) PizzaBuilder {
	b.toppings = toppings
	return b
}

func (b *pizzaBuilder) SetSize(size int) PizzaBuilder {
	b.size = size
	return b
}

func (b *pizzaBuilder) Build() (product.Pizza, error) {
	if err := validatePizza(b.crust, b.size, b.toppings); err != nil {
		return product.Pizza{}, err
	}
	return product.Pizza{
		Crust:    b.crust,
		Toppings: b.toppings,
		Size:     b.size,
	}, nil
}
