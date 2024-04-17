package config

// PizzaRecipe is a predefined pizza configuration.
type PizzaRecipe struct {
	Name     string
	Crust    string
	Toppings []string
	Size     int
}

// StandardRecipes returns a list of standard pizza recipes.
func StandardRecipes() []PizzaRecipe {
	return []PizzaRecipe{
		{"Margherita", "thin", []string{"mozzarella", "tomato", "basil"}, 12},
		{"Pepperoni", "traditional", []string{"pepperoni", "mozzarella"}, 12},
	}
}

// GourmetRecipes returns a list of gourmet pizza recipes.
func GourmetRecipes() []PizzaRecipe {
	return []PizzaRecipe{
		{"Truffle", "wood-fired", []string{"truffle oil", "mushrooms", "parmesan"}, 14},
		{"Seafood", "thin", []string{"shrimp", "squid", "garlic"}, 14},
	}
}
