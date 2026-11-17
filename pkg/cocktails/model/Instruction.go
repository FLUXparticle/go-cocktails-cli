package model

type Instruction struct {
	CocktailID   uint        `json:"-"`
	AmountCL     int         `json:"amount"`
	IngredientID uint        `json:"ingredient_id,omitempty"`
	Ingredient   *Ingredient `json:"ingredient"`
}
