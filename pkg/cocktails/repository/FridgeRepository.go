package repository

type FridgeRepository interface {
	// Ingredients liefert die vorhandenen Zutaten-IDs. Der Wert true steht für eine vorhandene Zutat.
	Ingredients() (map[uint]bool, error)
	// SetFridgeIngredient setzt für die Zutaten-ID id, ob sie im Kühlschrank liegt. present gibt den gewünschten Zustand an.
	SetFridgeIngredient(id uint, present bool) error
}
