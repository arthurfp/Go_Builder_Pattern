package builder

import "builder-pattern-go/pkg/product"

// pizzaBuilder builds a customizable pizza.
type pizzaBuilder struct {
	crust    string
	toppings []string
	size     int
}

// NewPizzaBuilder initializes a new builder for pizza creation.
func NewPizzaBuilder() PizzaBuilder {
	return &pizzaBuilder{}
}

// SetCrust chooses the crust type for the pizza.
func (b *pizzaBuilder) SetCrust(crust string) PizzaBuilder {
	b.crust = crust
	return b
}

// SetToppings adds toppings to the pizza. Multiple toppings are allowed.
func (b *pizzaBuilder) SetToppings(toppings []string) PizzaBuilder {
	b.toppings = toppings
	return b
}

// SetSize defines the size of the pizza (in inches).
func (b *pizzaBuilder) SetSize(size int) PizzaBuilder {
	b.size = size
	return b
}

// Build completes the pizza construction and returns the finished product.
func (b *pizzaBuilder) Build() product.Pizza {
	return product.Pizza{
		Crust:    b.crust,
		Toppings: b.toppings,
		Size:     b.size,
	}
}
