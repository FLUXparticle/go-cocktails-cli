package service

import (
	"sort"

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
	stock, err := s.repo.Ingredients()
	if err != nil {
		return nil, err
	}
	ingredients, err := s.serv.GetIngredients()
	if err != nil {
		return nil, err
	}

	rows := make([]*model.FridgeIngredient, 0, len(ingredients))
	for _, ingredient := range ingredients {
		rows = append(rows, &model.FridgeIngredient{Ingredient: *ingredient, InFridge: stock[ingredient.IngredientID]})
	}
	return rows, nil
}

func (s *FridgeService) SetFridgeIngredient(id uint, present bool) error {
	return s.repo.SetFridgeIngredient(id, present)
}

func (s *FridgeService) GetPossibleCocktails() ([]*model.Cocktail, error) {
	stock, err := s.repo.Ingredients()
	if err != nil {
		return nil, err
	}
	possible := make([]*model.Cocktail, 0)
	checked := make(map[uint]bool)
	for id, present := range stock {
		if !present {
			continue
		}
		_, candidates, err := s.serv.GetCocktailsByIngredient(id)
		if err != nil {
			return nil, err
		}
		for _, cocktail := range candidates {
			if checked[cocktail.CocktailID] {
				continue
			}
			checked[cocktail.CocktailID] = true
			recipe, err := s.serv.GetCocktail(cocktail.CocktailID)
			if err != nil {
				return nil, err
			}
			if canMix(recipe, stock) {
				possible = append(possible, cocktail)
			}
		}
	}
	sort.Slice(possible, func(i int, j int) bool { return possible[i].CocktailID < possible[j].CocktailID })
	return possible, nil
}

func canMix(recipe *model.Cocktail, stock map[uint]bool) bool {
	for _, instruction := range recipe.Instructions {
		if instruction.Ingredient == nil || !stock[instruction.Ingredient.IngredientID] {
			return false
		}
	}
	return len(recipe.Instructions) != 0
}
