package memory

import (
	"sync"
	"testing"
)

// TestFridgeConcurrentAccess prüft zwei gleichzeitige Schreibzugriffe auf denselben Kühlschrank.
func TestFridgeConcurrentAccess(t *testing.T) {
	fridge := NewFridgeRepository()
	const ingredientID = 31
	const count = 10_000

	var workers sync.WaitGroup
	workers.Go(func() {
		for range count {
			if err := fridge.SetFridgeIngredient(ingredientID, true); err != nil {
				t.Error(err)
				return
			}
		}
	})
	workers.Go(func() {
		for range count {
			if err := fridge.SetFridgeIngredient(ingredientID, false); err != nil {
				t.Error(err)
				return
			}
		}
	})
	workers.Wait()
}
