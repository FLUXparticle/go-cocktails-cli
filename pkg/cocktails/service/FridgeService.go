package service

import (
	"go-cocktails-cli/pkg/cocktails/model"
	"go-cocktails-cli/pkg/cocktails/repository"
)

type FridgeService struct {
	repo repository.FridgeRepository
	serv *CocktailService
}

// NewFridgeService erstellt einen Service. repo speichert den Kühlschrankbestand und serv liefert Zutaten und Rezepte.
func NewFridgeService(repo repository.FridgeRepository, serv *CocktailService) *FridgeService {
	return &FridgeService{repo: repo, serv: serv}
}

// GetFridge liefert alle Katalogzutaten mit der Angabe, ob sie im Kühlschrank liegen.
func (s *FridgeService) GetFridge() ([]*model.FridgeIngredient, error) {
	// TODO Aufgabe 5: Hole den Kühlschrankbestand.
	panic("Aufgabe 5")
}

// SetFridgeIngredient setzt den Bestandseintrag für die Zutaten-ID id. present gibt an, ob die Zutat vorhanden sein soll.
func (s *FridgeService) SetFridgeIngredient(id uint, present bool) error {
	// TODO Aufgabe 5: Ändere eine Zutat im Bestand.
	panic("Aufgabe 5")
}

// GetPossibleCocktails liefert die Cocktails, für die alle benötigten Zutaten im Kühlschrank liegen.
func (s *FridgeService) GetPossibleCocktails() ([]*model.Cocktail, error) {
	// TODO Aufgabe 5: Ermittle passende Cocktails.
	panic("Aufgabe 5")
}

// canMix prüft, ob alle Zutaten des recipe im Bestand stock liegen.
// In stock bedeutet true, dass die Zutat mit der jeweiligen ID vorhanden ist.
func canMix(recipe *model.Cocktail, stock map[uint]bool) bool {
	// TODO Aufgabe 5: Prüfe die Zutaten eines Rezepts.
	panic("Aufgabe 5")
}
