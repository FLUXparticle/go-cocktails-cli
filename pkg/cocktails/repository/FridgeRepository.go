package repository

type FridgeRepository interface {
	Ingredients() (map[uint]bool, error)
	SetFridgeIngredient(id uint, present bool) error
}
