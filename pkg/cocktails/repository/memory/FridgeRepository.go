package memory

import "maps"

type FridgeRepository struct {
	ingredients map[uint]bool
}

// NewFridgeRepository legt einen leeren Kühlschrankbestand im Speicher an.
func NewFridgeRepository() *FridgeRepository {
	return &FridgeRepository{ingredients: make(map[uint]bool)}
}

// SetFridgeIngredient setzt den Bestandseintrag für die Zutaten-ID id. present gibt an, ob die Zutat im Kühlschrank liegt.
func (r *FridgeRepository) SetFridgeIngredient(id uint, present bool) error {
	// TODO Aufgabe 7: Schütze den Schreibzugriff mit einem Mutex.
	r.ingredients[id] = present
	return nil
}

// Ingredients liefert eine Kopie der Zutaten-Map, die der Aufrufer unabhängig ändern kann.
func (r *FridgeRepository) Ingredients() (map[uint]bool, error) {
	// TODO Aufgabe 7: Schütze das Kopieren der Map mit dem Leselock.
	return maps.Clone(r.ingredients), nil
}
