package main

import (
	"flag"
	"fmt"
	"os"

	"go-cocktails-cli/cmd/internal/console"
	"go-cocktails-cli/pkg/cocktails/repository/remote"
	"go-cocktails-cli/pkg/cocktails/service"
)

func main() {
	add := flag.Uint("add", 0, "ID der hinzuzufügenden Zutat")
	remove := flag.Uint("remove", 0, "ID der zu entfernenden Zutat")
	display := flag.Bool("display", false, "Inhalt des Kühlschranks anzeigen")
	possible := flag.Bool("possible", false, "Mögliche Cocktails anzeigen")
	flag.Parse()
	if flag.NArg() != 0 {
		flag.Usage()
		os.Exit(2)
	}
	if err := run(*add, *remove, *display, *possible); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(add uint, remove uint, display bool, possible bool) error {
	repo := remote.NewRemoteRepository()
	serv := service.NewCocktailService(repo)
	fridge := service.NewFridgeService(repo, serv)

	if add != 0 {
		if err := fridge.SetFridgeIngredient(add, true); err != nil {
			return err
		}
	}

	if remove != 0 {
		if err := fridge.SetFridgeIngredient(remove, false); err != nil {
			return err
		}
	}

	if display {
		rows, err := fridge.GetFridge()
		if err != nil {
			return err
		}
		for _, row := range rows {
			if row.InFridge {
				fmt.Printf("%d: %s\n", row.IngredientID, row.Name)
			}
		}
	}

	if possible {
		rows, err := fridge.GetPossibleCocktails()
		if err != nil {
			return err
		}
		console.PrintCocktails(rows)
	}

	return nil
}
