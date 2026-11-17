package repository

import (
	"errors"

	"go-cocktails-cli/pkg/cocktails/model"
)

var ErrNotFound = errors.New("resource not found")

type CocktailRepository interface {
	GetCocktails() ([]*model.Cocktail, error)
	GetCocktail(id uint) (*model.Cocktail, error)
	GetIngredients() ([]*model.Ingredient, error)
	FindCocktailsByName(query string) ([]*model.Cocktail, error)
	FindIngredientsByName(query string) ([]*model.Ingredient, error)
	GetCocktailsByIngredient(id uint) (string, []*model.Cocktail, error)
}
