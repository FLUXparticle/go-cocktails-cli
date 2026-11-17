package model

type FridgeIngredient struct {
	Ingredient
	InFridge bool `json:"inFridge"`
}
