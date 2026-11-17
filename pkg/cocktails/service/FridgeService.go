package service

import (
	"go-cocktails-cli/pkg/cocktails/model"
	"go-cocktails-cli/pkg/cocktails/repository"
)

type FridgeService struct {
	repo repository.FridgeRepository
	serv *CocktailService
}

func NewFridgeService(repo repository.FridgeRepository, serv *CocktailService) *FridgeService {
	return &FridgeService{repo: repo, serv: serv}
}

func (s *FridgeService) GetFridge() ([]*model.FridgeIngredient, error) {
	// TODO Aufgabe 5: Hole den Kühlschrankbestand.
	panic("Aufgabe 5")
}

func (s *FridgeService) SetFridgeIngredient(id uint, present bool) error {
	// TODO Aufgabe 5: Ändere eine Zutat im Bestand.
	panic("Aufgabe 5")
}

func (s *FridgeService) GetPossibleCocktails() ([]*model.Cocktail, error) {
	// TODO Aufgabe 5: Ermittle passende Cocktails.
	panic("Aufgabe 5")
}

func canMix(recipe *model.Cocktail, stock map[uint]bool) bool {
	// TODO Aufgabe 5: Prüfe die Zutaten eines Rezepts.
	panic("Aufgabe 5")
}
