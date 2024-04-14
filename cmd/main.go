package main

import (
	"builder-pattern-go/pkg/builder"
	"fmt"
)

func main() {
	fmt.Println("Demonstrating the Builder Pattern in Go")

	// Initialize a new pizza builder
	pizzaBuilder := builder.NewPizzaBuilder()

	// Configure and build a custom pizza
	pizza := pizzaBuilder.
		SetCrust("thin").
		SetToppings([]string{"cheese", "tomatoes"}).
		SetSize(12).
		Build()

	// Display the details of the constructed pizza
	fmt.Printf("Constructed Pizza: Crust - %s, Toppings - %v, Size - %d inches\n", pizza.Crust, pizza.Toppings, pizza.Size)
}
