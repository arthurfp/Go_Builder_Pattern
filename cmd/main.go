package main

import (
	"builder-pattern-go/pkg/builder"
	"builder-pattern-go/pkg/config"
	"fmt"
)

func main() {
	fmt.Println("Demonstrating the Builder Pattern in Go with various builders and recipes")

	// Initialize a standard pizza builder
	pizzaBuilder := builder.NewPizzaBuilder()
	standardRecipes := config.StandardRecipes()

	for _, recipe := range standardRecipes {
		pizza := pizzaBuilder.SetCrust(recipe.Crust).
			SetToppings(recipe.Toppings).
			SetSize(recipe.Size).
			Build()
		fmt.Printf("Standard Pizza - %s: Crust - %s, Toppings - %v, Size - %d inches\n", recipe.Name, pizza.Crust, pizza.Toppings, pizza.Size)
	}

	// Initialize a gourmet pizza builder
	gourmetBuilder := builder.NewGourmetPizzaBuilder()
	gourmetRecipes := config.GourmetRecipes()

	for _, recipe := range gourmetRecipes {
		gourmetPizza := gourmetBuilder.SetCrust(recipe.Crust).
			SetToppings(recipe.Toppings).
			SetSize(recipe.Size).
			Build()
		fmt.Printf("Gourmet Pizza - %s: Crust - %s, Toppings - %v, Size - %d inches\n", recipe.Name, gourmetPizza.Crust, gourmetPizza.Toppings, gourmetPizza.Size)
	}
}
