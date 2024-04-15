package main

import (
	"builder-pattern-go/pkg/builder"
	"fmt"
)

func main() {
	fmt.Println("Demonstrating the Builder Pattern in Go with various builders")

	// Initialize a standard pizza builder
	pizzaBuilder := builder.NewPizzaBuilder()
	pizza := pizzaBuilder.SetCrust("thin").SetToppings([]string{"cheese", "tomatoes"}).SetSize(12).Build()
	fmt.Printf("Standard Pizza: Crust - %s, Toppings - %v, Size - %d inches\n", pizza.Crust, pizza.Toppings, pizza.Size)

	// Initialize a gourmet pizza builder
	gourmetBuilder := builder.NewGourmetPizzaBuilder()
	gourmetPizza := gourmetBuilder.SetCrust("wood-fired").SetToppings([]string{"goat cheese", "sun-dried tomatoes", "arugula"}).SetSize(14).Build()
	fmt.Printf("Gourmet Pizza: Crust - %s, Toppings - %v, Size - %d inches\n", gourmetPizza.Crust, gourmetPizza.Toppings, gourmetPizza.Size)
}
