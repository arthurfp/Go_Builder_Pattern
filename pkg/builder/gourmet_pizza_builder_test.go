package builder

import (
	"builder-pattern-go/pkg/product"
	"reflect"
	"testing"
)

// TestGourmetPizzaBuilder ensures that the gourmet pizza builder constructs pizzas correctly.
func TestGourmetPizzaBuilder(t *testing.T) {
	builder := NewGourmetPizzaBuilder()
	pizza, err := builder.SetCrust("wood-fired").SetToppings([]string{"truffle oil", "mushrooms"}).SetSize(14).Build()
	if err != nil {
		t.Fatalf("Failed to build gourmet pizza: %s", err)
	}

	expected := product.Pizza{
		Crust:    "wood-fired",
		Toppings: []string{"truffle oil", "mushrooms"},
		Size:     14,
	}

	if !reflect.DeepEqual(pizza, expected) {
		t.Errorf("Expected pizza %+v, got %+v", expected, pizza)
	}
}

// TestGourmetPizzaBuilderError checks that the builder correctly handles invalid configurations.
func TestGourmetPizzaBuilderError(t *testing.T) {
	builder := NewGourmetPizzaBuilder()
	_, err := builder.SetCrust("wood-fired").SetToppings([]string{}).SetSize(20).Build()
	if err == nil {
		t.Error("Expected error for invalid gourmet pizza size and lack of toppings, got nil")
	}
}
