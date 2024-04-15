package builder

import "builder-pattern-go/pkg/product"

// gourmetPizzaBuilder builds gourmet-style pizzas.
type gourmetPizzaBuilder struct {
	crust    string
	toppings []string
	size     int
}

// NewGourmetPizzaBuilder creates a new builder for gourmet pizza creation.
func NewGourmetPizzaBuilder() PizzaBuilder {
	return &gourmetPizzaBuilder{}
}

// SetCrust sets the crust type for the gourmet pizza.
func (b *gourmetPizzaBuilder) SetCrust(crust string) PizzaBuilder {
	b.crust = crust
	return b
}

// SetToppings adds gourmet toppings to the pizza.
func (b *gourmetPizzaBuilder) SetToppings(toppings []string) PizzaBuilder {
	b.toppings = toppings
	return b
}

// SetSize defines the size of the gourmet pizza.
func (b *gourmetPizzaBuilder) SetSize(size int) PizzaBuilder {
	b.size = size
	return b
}

// Build completes the gourmet pizza construction and returns the finished product.
func (b *gourmetPizzaBuilder) Build() product.Pizza {
	return product.Pizza{
		Crust:    b.crust,
		Toppings: b.toppings,
		Size:     b.size,
	}
}
