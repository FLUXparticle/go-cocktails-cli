package memory

import "maps"

type FridgeRepository struct {
	ingredients map[uint]bool
}

func NewFridgeRepository() *FridgeRepository {
	return &FridgeRepository{ingredients: make(map[uint]bool)}
}

func (r *FridgeRepository) SetFridgeIngredient(id uint, present bool) error {
	// TODO Aufgabe 7: Schütze den Schreibzugriff mit einem Mutex.
	r.ingredients[id] = present
	return nil
}

// Ingredients returns a snapshot that callers may use after the lock is released.
func (r *FridgeRepository) Ingredients() (map[uint]bool, error) {
	// TODO Aufgabe 7: Schütze das Kopieren der Map mit dem Leselock.
	return maps.Clone(r.ingredients), nil
}
