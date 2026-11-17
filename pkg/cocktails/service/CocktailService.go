package service

import (
	"go-cocktails-cli/pkg/cocktails/model"
	"go-cocktails-cli/pkg/cocktails/repository"
	"golang.org/x/sync/errgroup"
	"sort"
	"strings"
)

type CocktailService struct {
	repo repository.CocktailRepository
}

func NewCocktailService(repo repository.CocktailRepository) *CocktailService {
	return &CocktailService{repo: repo}
}

func (s *CocktailService) MilkAmount() (int, error) {
	cocktails, err := s.GetCocktails()
	if err != nil {
		return 0, err
	}

	total := 0
	for _, cocktail := range cocktails {
		if !strings.Contains(cocktail.Name, "Milk") {
			continue
		}
		recipe, err := s.GetCocktail(cocktail.CocktailID)
		if err != nil {
			return 0, err
		}
		for _, instruction := range recipe.Instructions {
			if instruction.Ingredient != nil && instruction.Ingredient.Name == "Milch" {
				total += instruction.AmountCL
			}
		}
	}
	return total, nil
}

func (s *CocktailService) MilkAmountParallel() (int, error) {
	// TODO Aufgabe 2: Berechne die Milchmenge mit parallelen Rezeptabrufen.
	panic("Aufgabe 2")
}

func (s *CocktailService) MilkAmountErrGroup(group *errgroup.Group) (int, error) {
	// TODO Aufgabe 3: Summiere parallele Abrufe und brich bei einem Fehler ab.
	panic("Aufgabe 3")
}

func (s *CocktailService) GetCocktails() ([]*model.Cocktail, error) {
	return s.repo.GetCocktails()
}

func (s *CocktailService) GetCocktail(id uint) (*model.Cocktail, error) {
	return s.repo.GetCocktail(id)
}

func (s *CocktailService) GetIngredients() ([]*model.Ingredient, error) {
	return s.repo.GetIngredients()
}

func (s *CocktailService) GetCocktailsByIngredient(id uint) (string, []*model.Cocktail, error) {
	return s.repo.GetCocktailsByIngredient(id)
}

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
