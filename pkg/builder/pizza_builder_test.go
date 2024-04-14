package builder

import (
	"builder-pattern-go/pkg/product"
	"reflect"
	"testing"
)

// TestPizzaBuilder ensures that the builder correctly sets properties and builds the pizza.
func TestPizzaBuilder(t *testing.T) {
	builder := NewPizzaBuilder()
	pizza := builder.SetCrust("thick").SetToppings([]string{"pepperoni", "mushrooms"}).SetSize(16).Build()

	expected := product.Pizza{
		Crust:    "thick",
		Toppings: []string{"pepperoni", "mushrooms"},
		Size:     16,
	}

	if !reflect.DeepEqual(pizza, expected) {
		t.Errorf("Expected pizza %+v, got %+v", expected, pizza)
	}
}
