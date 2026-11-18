package service

import (
	"go-cocktails-cli/pkg/cocktails/model"
	"go-cocktails-cli/pkg/cocktails/repository"
	"golang.org/x/sync/errgroup"
	"sort"
)

type CocktailService struct {
	repo repository.CocktailRepository
}

// NewCocktailService erstellt einen Service. repo liefert die Cocktail- und Zutatendaten.
func NewCocktailService(repo repository.CocktailRepository) *CocktailService {
	return &CocktailService{repo: repo}
}

// MilkAmount berechnet die gesamte Milchmenge in Zentilitern für Cocktails mit Milk im Namen. Die Detailabrufe laufen nacheinander.
func (s *CocktailService) MilkAmount() (int, error) {
	// TODO Aufgabe 1: Berechne die Milchmenge der passenden Cocktails.
	panic("Aufgabe 1")
}

// MilkAmountParallel berechnet dieselbe Milchmenge mit parallelen Detailabrufen.
func (s *CocktailService) MilkAmountParallel() (int, error) {
	// TODO Aufgabe 2: Berechne die Milchmenge mit parallelen Rezeptabrufen.
	panic("Aufgabe 2")
}

// MilkAmountErrGroup startet die parallelen Detailabrufe in group und gibt deren ersten Fehler zurück.
// group muss bei einem Fehler den Kontext abbrechen, den auch das Remote-Repository verwendet.
// Für die Übung wird jede gefundene Cocktail-ID vor dem Detailabruf um 45 erhöht.
func (s *CocktailService) MilkAmountErrGroup(group *errgroup.Group) (int, error) {
	// TODO Aufgabe 3: Summiere parallele Abrufe und brich bei einem Fehler ab.
	panic("Aufgabe 3")
}

// GetCocktails liefert die Cocktail-Liste mit IDs und Namen.
func (s *CocktailService) GetCocktails() ([]*model.Cocktail, error) {
	return s.repo.GetCocktails()
}

// GetCocktail liefert das Rezept für die Cocktail-ID id samt Anweisungen und Zutaten.
func (s *CocktailService) GetCocktail(id uint) (*model.Cocktail, error) {
	return s.repo.GetCocktail(id)
}

// GetIngredients liefert die Zutatenliste mit IDs und Namen.
func (s *CocktailService) GetIngredients() ([]*model.Ingredient, error) {
	return s.repo.GetIngredients()
}

// GetCocktailsByIngredient liefert den Namen der Zutat mit der ID id und die Cocktails, die sie enthalten.
func (s *CocktailService) GetCocktailsByIngredient(id uint) (string, []*model.Cocktail, error) {
	return s.repo.GetCocktailsByIngredient(id)
}

// Search liefert Cocktails, deren Name oder Zutatenname den Suchbegriff query enthält.
func (s *CocktailService) Search(query string) ([]*model.Cocktail, error) {
	cocktails, err := s.repo.FindCocktailsByName(query)
	if err != nil {
		return nil, err
	}
	// An empty substring already matches every cocktail.
	if query == "" {
		return cocktails, nil
	}
	ingredients, err := s.repo.FindIngredientsByName(query)
	if err != nil {
		return nil, err
	}

	seen := make(map[uint]bool)
	for _, cocktail := range cocktails {
		seen[cocktail.CocktailID] = true
	}
	for _, ingredient := range ingredients {
		_, matches, err := s.repo.GetCocktailsByIngredient(ingredient.IngredientID)
		if err != nil {
			return nil, err
		}
		for _, cocktail := range matches {
			if !seen[cocktail.CocktailID] {
				seen[cocktail.CocktailID] = true
				cocktails = append(cocktails, cocktail)
			}
		}
	}
	sort.Slice(cocktails, func(i int, j int) bool { return cocktails[i].CocktailID < cocktails[j].CocktailID })
	return cocktails, nil
}
