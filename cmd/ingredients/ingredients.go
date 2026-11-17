package main

import (
	"flag"
	"fmt"
	"os"

	"go-cocktails-cli/cmd/internal/console"
	"go-cocktails-cli/pkg/cocktails/repository/database"
	"go-cocktails-cli/pkg/cocktails/service"
)

func main() {
	var id uint
	flag.UintVar(&id, "id", 0, "ID der Zutat")
	flag.UintVar(&id, "i", 0, "Kurzform von -id")
	flag.Parse()
	if flag.NArg() != 0 {
		fmt.Fprintln(os.Stderr, "ingredients erwartet keine Positionsargumente")
		os.Exit(1)
	}
	if err := run(id); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(id uint) error {
	repo, err := database.NewDatabaseRepository()
	if err != nil {
		return err
	}
	defer repo.Close()

	svc := service.NewCocktailService(repo)
	if id == 0 {
		ingredients, err := svc.GetIngredients()
		if err != nil {
			return err
		}
		for _, ingredient := range ingredients {
			fmt.Printf("%d: %s\n", ingredient.IngredientID, ingredient.Name)
		}
		return nil
	}
	name, cocktails, err := svc.GetCocktailsByIngredient(id)
	if err != nil {
		return err
	}
	fmt.Fprintln(os.Stdout, name)
	console.PrintCocktails(cocktails)
	return nil
}
