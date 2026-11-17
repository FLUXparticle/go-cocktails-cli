package memory

import (
	"maps"
	"sync"
)

type FridgeRepository struct {
	mu          sync.RWMutex
	ingredients map[uint]bool
}

func NewFridgeRepository() *FridgeRepository {
	return &FridgeRepository{ingredients: make(map[uint]bool)}
}

func (r *FridgeRepository) SetFridgeIngredient(id uint, present bool) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.ingredients[id] = present
	return nil
}

// Ingredients returns a snapshot that callers may use after the lock is released.
func (r *FridgeRepository) Ingredients() (map[uint]bool, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return maps.Clone(r.ingredients), nil
}
