package builder

import (
	"builder-pattern-go/pkg/product"
)

// PizzaBuilder interface defines the steps to construct the pizza product.
type PizzaBuilder interface {
	SetCrust(crust string) PizzaBuilder
	SetToppings(toppings []string) PizzaBuilder
	SetSize(size int) PizzaBuilder
	Build() product.Pizza
}
