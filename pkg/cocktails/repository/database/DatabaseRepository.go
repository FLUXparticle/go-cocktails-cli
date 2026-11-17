package database

import (
	"database/sql"
	"errors"
	"fmt"
	"net/url"
	"path/filepath"
	"strings"

	_ "github.com/mattn/go-sqlite3"

	"go-cocktails-cli/assets"
	"go-cocktails-cli/pkg/cocktails/model"
	"go-cocktails-cli/pkg/cocktails/repository"
)

const databaseFilename = "cocktails.db"

type DatabaseRepository struct {
	db *sql.DB
}

func NewDatabaseRepository() (*DatabaseRepository, error) {
	dir, err := assets.Dir()
	if err != nil {
		return nil, err
	}
	path := filepath.Join(dir, databaseFilename)
	dsn := url.URL{Scheme: "file", Path: filepath.ToSlash(path), RawQuery: "mode=ro"}
	db, err := sql.Open("sqlite3", dsn.String())
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("Katalog %s öffnen: %w", path, err)
	}
	return &DatabaseRepository{db: db}, nil
}

func (r *DatabaseRepository) Close() error {
	return r.db.Close()
}

func lookupError(err error) error {
	if errors.Is(err, sql.ErrNoRows) {
		return repository.ErrNotFound
	}
	return err
}

func (r *DatabaseRepository) listCocktails(query string, args ...any) ([]*model.Cocktail, error) {
	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	result := make([]*model.Cocktail, 0)
	for rows.Next() {
		var cocktail model.Cocktail
		if err := rows.Scan(&cocktail.CocktailID, &cocktail.Name); err != nil {
			return nil, err
		}
		result = append(result, &cocktail)
	}
	return result, rows.Err()
}

func (r *DatabaseRepository) listIngredients(query string, args ...any) ([]*model.Ingredient, error) {
	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	result := make([]*model.Ingredient, 0)
	for rows.Next() {
		var ingredient model.Ingredient
		if err := rows.Scan(&ingredient.IngredientID, &ingredient.Name); err != nil {
			return nil, err
		}
		result = append(result, &ingredient)
	}
	return result, rows.Err()
}

func (r *DatabaseRepository) GetCocktails() ([]*model.Cocktail, error) {
	return r.listCocktails("SELECT cocktail_id, name FROM cocktails ORDER BY cocktail_id")
}

func (r *DatabaseRepository) GetCocktail(id uint) (*model.Cocktail, error) {
	cocktail := &model.Cocktail{CocktailID: id}
	err := r.db.QueryRow("SELECT name FROM cocktails WHERE cocktail_id = ?", id).Scan(&cocktail.Name)
	if err != nil {
		return nil, lookupError(err)
	}
	rows, err := r.db.Query(`
  SELECT i.amount_cl, i.ingredient_id, g.name
  FROM instructions i LEFT JOIN ingredients g ON g.ingredient_id = i.ingredient_id
  WHERE i.cocktail_id = ? ORDER BY i.rowid`, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		instruction := &model.Instruction{CocktailID: id}
		var name sql.NullString
		if err := rows.Scan(&instruction.AmountCL, &instruction.IngredientID, &name); err != nil {
			return nil, err
		}
		if name.Valid {
			instruction.Ingredient = &model.Ingredient{IngredientID: instruction.IngredientID, Name: name.String}
		}
		cocktail.Instructions = append(cocktail.Instructions, instruction)
	}
	return cocktail, rows.Err()
}

func (r *DatabaseRepository) GetIngredients() ([]*model.Ingredient, error) {
	return r.listIngredients("SELECT ingredient_id, name FROM ingredients ORDER BY ingredient_id")
}

func (r *DatabaseRepository) GetCocktailsByIngredient(id uint) (string, []*model.Cocktail, error) {
	var name string
	if err := r.db.QueryRow("SELECT name FROM ingredients WHERE ingredient_id = ?", id).Scan(&name); err != nil {
		return "", nil, lookupError(err)
	}
	cocktails, err := r.listCocktails(`
  SELECT cocktail_id, name FROM cocktails
  WHERE cocktail_id IN (SELECT cocktail_id FROM instructions WHERE ingredient_id = ?)
  ORDER BY cocktail_id`, id)
	return name, cocktails, err
}

func namePattern(query string) string {
	return "%" + strings.NewReplacer("\\", "\\\\", "%", "\\%", "_", "\\_").Replace(query) + "%"
}

func (r *DatabaseRepository) FindCocktailsByName(query string) ([]*model.Cocktail, error) {
	return r.listCocktails(`SELECT cocktail_id, name FROM cocktails WHERE name LIKE ? ESCAPE '\' ORDER BY cocktail_id`, namePattern(query))
}

func (r *DatabaseRepository) FindIngredientsByName(query string) ([]*model.Ingredient, error) {
	return r.listIngredients(`SELECT ingredient_id, name FROM ingredients WHERE name LIKE ? ESCAPE '\' ORDER BY ingredient_id`, namePattern(query))
}
