package builder

import (
	"builder-pattern-go/pkg/product"
	"reflect"
	"testing"
)

// TestPizzaBuilder ensures that the builder correctly sets properties and builds the pizza.
func TestPizzaBuilder(t *testing.T) {
	builder := NewPizzaBuilder()
	pizza, err := builder.SetCrust("thin").SetToppings([]string{"cheese", "basil"}).SetSize(12).Build()
	if err != nil {
		t.Fatalf("Failed to build pizza: %s", err)
	}

	expected := product.Pizza{
		Crust:    "thin",
		Toppings: []string{"cheese", "basil"},
		Size:     12,
	}

	if !reflect.DeepEqual(pizza, expected) {
		t.Errorf("Expected pizza %+v, got %+v", expected, pizza)
	}
}

// TestPizzaBuilderError verifies that the builder returns an error for invalid configurations.
func TestPizzaBuilderError(t *testing.T) {
	builder := NewPizzaBuilder()
	_, err := builder.SetCrust("thin").SetToppings([]string{}).SetSize(5).Build()
	if err == nil {
		t.Error("Expected error for invalid pizza size and toppings, got nil")
	}
}
