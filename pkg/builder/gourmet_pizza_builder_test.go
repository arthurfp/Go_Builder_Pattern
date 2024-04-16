package builder

import (
	"builder-pattern-go/pkg/product"
	"reflect"
	"testing"
)

// TestGourmetPizzaBuilder verifies the gourmet pizza builder constructs pizzas as expected.
func TestGourmetPizzaBuilder(t *testing.T) {
	builder := NewGourmetPizzaBuilder()
	pizza := builder.SetCrust("wood-fired").SetToppings([]string{"goat cheese", "sun-dried tomatoes", "arugula"}).SetSize(14).Build()

	expected := product.Pizza{
		Crust:    "wood-fired",
		Toppings: []string{"goat cheese", "sun-dried tomatoes", "arugula"},
		Size:     14,
	}

	if !reflect.DeepEqual(pizza, expected) {
		t.Errorf("Expected pizza %+v, got %+v", expected, pizza)
	}
}
