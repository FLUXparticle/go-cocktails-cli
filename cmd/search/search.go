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
	flag.Usage = func() {
		fmt.Fprintln(flag.CommandLine.Output(), "Aufruf: search [Optionen] SUCHBEGRIFF")
		flag.PrintDefaults()
	}
	flag.Parse()
	if flag.NArg() != 1 {
		flag.Usage()
		os.Exit(2)
	}
	if err := run(flag.Arg(0)); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

// run zeigt Cocktails an, deren Name oder Zutatenname den Suchbegriff query enthält.
func run(query string) error {
	repo, err := database.NewDatabaseRepository()
	if err != nil {
		return err
	}
	defer repo.Close()
	cocktails, err := service.NewCocktailService(repo).Search(query)
	if err != nil {
		return err
	}
	console.PrintCocktails(cocktails)
	return nil
}
