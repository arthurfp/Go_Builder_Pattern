package builder

import (
	"builder-pattern-go/pkg/product"
	"errors"
)

// PizzaBuilder interface defines the steps to construct the pizza product.
type PizzaBuilder interface {
	SetCrust(crust string) PizzaBuilder
	SetToppings(toppings []string) PizzaBuilder
	SetSize(size int) PizzaBuilder
	Build() (product.Pizza, error)
}

// Validate size and other parameters.
func validatePizza(crust string, size int, toppings []string) error {
	if size < 8 || size > 16 {
		return errors.New("invalid size: size must be between 8 and 16 inches")
	}
	if len(toppings) == 0 {
		return errors.New("invalid configuration: at least one topping is required")
	}
	return nil
}
