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
	flag.UintVar(&id, "id", 0, "ID des Cocktails")
	flag.UintVar(&id, "i", 0, "Kurzform von -id")
	flag.Parse()
	if flag.NArg() != 0 {
		fmt.Fprintln(os.Stderr, "cocktail erwartet keine Positionsargumente")
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
	if id != 0 {
		cocktail, err := svc.GetCocktail(id)
		if err != nil {
			return err
		}
		console.PrintRecipe(cocktail)
		return nil
	}
	cocktails, err := svc.GetCocktails()
	if err != nil {
		return err
	}
	console.PrintCocktails(cocktails)
	return nil
}
