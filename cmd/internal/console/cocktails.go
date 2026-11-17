package console

import (
	"fmt"

	"go-cocktails-cli/pkg/cocktails/model"
)

func PrintCocktails(cocktails []*model.Cocktail) {
	for _, cocktail := range cocktails {
		fmt.Printf("%d: %s\n", cocktail.CocktailID, cocktail.Name)
	}
}

func PrintRecipe(cocktail *model.Cocktail) {
	fmt.Printf("%s (ID: %d)\n", cocktail.Name, cocktail.CocktailID)
	for _, instruction := range cocktail.Instructions {
		ingredient := fmt.Sprintf("Zutat %d", instruction.IngredientID)
		if instruction.Ingredient != nil {
			ingredient = instruction.Ingredient.Name
		}
		fmt.Printf("- %d cl %s\n", instruction.AmountCL, ingredient)
	}
}
