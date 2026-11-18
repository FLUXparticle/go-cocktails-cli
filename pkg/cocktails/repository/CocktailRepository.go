package repository

import (
	"errors"

	"go-cocktails-cli/pkg/cocktails/model"
)

var ErrNotFound = errors.New("resource not found")

type CocktailRepository interface {
	// GetCocktails liefert die Cocktail-Liste mit IDs und Namen.
	GetCocktails() ([]*model.Cocktail, error)
	// GetCocktail liefert das Rezept für die Cocktail-ID id samt Anweisungen und Zutaten.
	GetCocktail(id uint) (*model.Cocktail, error)
	// GetIngredients liefert die Zutatenliste mit IDs und Namen.
	GetIngredients() ([]*model.Ingredient, error)
	// FindCocktailsByName liefert die Cocktails, deren Name den Suchbegriff query enthält.
	FindCocktailsByName(query string) ([]*model.Cocktail, error)
	// FindIngredientsByName liefert die Zutaten, deren Name den Suchbegriff query enthält.
	FindIngredientsByName(query string) ([]*model.Ingredient, error)
	// GetCocktailsByIngredient liefert den Namen der Zutat mit der ID id und die Cocktails, die sie enthalten.
	GetCocktailsByIngredient(id uint) (string, []*model.Cocktail, error)
}
