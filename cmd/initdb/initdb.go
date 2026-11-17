package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	_ "github.com/mattn/go-sqlite3"
	"go-cocktails-cli/assets"
	"go-cocktails-cli/pkg/cocktails/model"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	dir, err := assets.Dir()
	if err != nil {
		return err
	}
	cocktails, ingredients, err := readDataset(filepath.Join(dir, "cocktails.txt"))
	if err != nil {
		return err
	}
	path := filepath.Join(dir, "cocktails.db")
	if err := writeDataset(path, cocktails, ingredients); err != nil {
		return err
	}
	fmt.Printf("%s: %d Cocktails, %d Zutaten\n", path, len(cocktails), len(ingredients))
	return nil
}

func readDataset(path string) ([]*model.Cocktail, []*model.Ingredient, error) {
	input, err := os.Open(path)
	if err != nil {
		return nil, nil, err
	}
	defer input.Close()

	var cocktails []*model.Cocktail
	var ingredients []*model.Ingredient
	ingredientIDs := make(map[string]uint)
	var current *model.Cocktail
	scanner := bufio.NewScanner(input)
	for lineNumber := 1; scanner.Scan(); lineNumber++ {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			current = nil
			continue
		}
		if current == nil {
			current = &model.Cocktail{CocktailID: uint(len(cocktails) + 1), Name: line}
			cocktails = append(cocktails, current)
			continue
		}

		amount := 0
		name := line
		if count, ingredient, ok := strings.Cut(line, "cl:"); ok {
			amount, err = strconv.Atoi(strings.TrimSpace(count))
			if err != nil || amount < 0 {
				return nil, nil, fmt.Errorf("%s:%d: ungültige Menge %q", path, lineNumber, count)
			}
			name = strings.TrimSpace(ingredient)
		}
		if name == "" {
			return nil, nil, fmt.Errorf("%s:%d: Zutatenname fehlt", path, lineNumber)
		}
		id, found := ingredientIDs[name]
		if !found {
			id = uint(len(ingredients) + 1)
			ingredientIDs[name] = id
			ingredients = append(ingredients, &model.Ingredient{IngredientID: id, Name: name})
		}
		current.Instructions = append(current.Instructions, &model.Instruction{
			CocktailID: current.CocktailID, IngredientID: id, AmountCL: amount,
		})
	}
	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}
	if len(cocktails) == 0 {
		return nil, nil, fmt.Errorf("%s: keine Cocktails gefunden", path)
	}
	return cocktails, ingredients, nil
}

func writeDataset(path string, cocktails []*model.Cocktail, ingredients []*model.Ingredient) error {
	temp, err := os.CreateTemp(filepath.Dir(path), ".cocktails-*.db")
	if err != nil {
		return err
	}
	defer os.Remove(temp.Name())
	if err := temp.Close(); err != nil {
		return err
	}

	dsn := url.URL{Scheme: "file", Path: filepath.ToSlash(temp.Name()), RawQuery: "_foreign_keys=on"}
	db, err := sql.Open("sqlite3", dsn.String())
	if err != nil {
		return err
	}
	defer db.Close()

	if err := importDataset(db, cocktails, ingredients); err != nil {
		return err
	}
	if err := db.Close(); err != nil {
		return err
	}
	return os.Rename(temp.Name(), path)
}

func importDataset(db *sql.DB, cocktails []*model.Cocktail, ingredients []*model.Ingredient) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if _, err := tx.Exec(`
  CREATE TABLE ingredients (
   ingredient_id INTEGER PRIMARY KEY,
   name TEXT NOT NULL
  );
  CREATE TABLE cocktails (
   cocktail_id INTEGER PRIMARY KEY,
   name TEXT NOT NULL
  );
  CREATE TABLE instructions (
   cocktail_id INTEGER NOT NULL REFERENCES cocktails(cocktail_id),
   amount_cl INTEGER NOT NULL,
   ingredient_id INTEGER NOT NULL REFERENCES ingredients(ingredient_id)
  );`); err != nil {
		return err
	}

	for _, ingredient := range ingredients {
		if _, err := tx.Exec("INSERT INTO ingredients (ingredient_id, name) VALUES (?, ?)", ingredient.IngredientID, ingredient.Name); err != nil {
			return err
		}
	}
	for _, cocktail := range cocktails {
		if _, err := tx.Exec("INSERT INTO cocktails (cocktail_id, name) VALUES (?, ?)", cocktail.CocktailID, cocktail.Name); err != nil {
			return err
		}
		for _, instruction := range cocktail.Instructions {
			if _, err := tx.Exec("INSERT INTO instructions (cocktail_id, amount_cl, ingredient_id) VALUES (?, ?, ?)", cocktail.CocktailID, instruction.AmountCL, instruction.IngredientID); err != nil {
				return err
			}
		}
	}
	return tx.Commit()
}
