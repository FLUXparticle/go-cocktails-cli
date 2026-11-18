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

// NewDatabaseRepository öffnet die mit initdb erzeugte SQLite-Datei schreibgeschützt.
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

// Close schließt den Datenbankzugriff des Repositorys.
func (r *DatabaseRepository) Close() error {
	return r.db.Close()
}

// lookupError übersetzt sql.ErrNoRows aus err in repository.ErrNotFound. Andere Fehler gibt die Funktion unverändert zurück.
func lookupError(err error) error {
	if errors.Is(err, sql.ErrNoRows) {
		return repository.ErrNotFound
	}
	return err
}

// listCocktails liest die Ergebniszeilen von query als Cocktails.
// args enthält die Werte für die SQL-Platzhalter in ihrer Reihenfolge.
// Die Abfrage muss zuerst die Cocktail-ID und dann den Namen liefern.
func (r *DatabaseRepository) listCocktails(query string, args ...any) ([]*model.Cocktail, error) {
	// TODO Aufgabe 6: Führe die übergebene SQL-Abfrage aus und lies die Cocktails ein.
	panic("Aufgabe 6")
}

// listIngredients liest die Ergebniszeilen von query als Zutaten.
// args enthält die Werte für die SQL-Platzhalter in ihrer Reihenfolge.
// Die Abfrage muss zuerst die Zutaten-ID und dann den Namen liefern.
func (r *DatabaseRepository) listIngredients(query string, args ...any) ([]*model.Ingredient, error) {
	// TODO Aufgabe 6: Führe die übergebene SQL-Abfrage aus und lies die Zutaten ein.
	panic("Aufgabe 6")
}

// GetCocktails liefert die Cocktail-Liste mit IDs und Namen.
func (r *DatabaseRepository) GetCocktails() ([]*model.Cocktail, error) {
	return r.listCocktails("SELECT cocktail_id, name FROM cocktails ORDER BY cocktail_id")
}

// GetCocktail liefert das Rezept für die Cocktail-ID id samt Anweisungen und Zutaten.
func (r *DatabaseRepository) GetCocktail(id uint) (*model.Cocktail, error) {
	// TODO Aufgabe 6: Lade Cocktail und Rezept und lies die Ergebnisse ein.
	// SELECT name FROM cocktails WHERE cocktail_id = ?
	// SELECT i.amount_cl, i.ingredient_id, g.name
	// FROM instructions i LEFT JOIN ingredients g ON g.ingredient_id = i.ingredient_id
	// WHERE i.cocktail_id = ? ORDER BY i.rowid
	panic("Aufgabe 6")
}

// GetIngredients liefert die Zutatenliste mit IDs und Namen.
func (r *DatabaseRepository) GetIngredients() ([]*model.Ingredient, error) {
	return r.listIngredients("SELECT ingredient_id, name FROM ingredients ORDER BY ingredient_id")
}

// GetCocktailsByIngredient liefert den Namen der Zutat mit der ID id und die Cocktails, die sie enthalten.
func (r *DatabaseRepository) GetCocktailsByIngredient(id uint) (string, []*model.Cocktail, error) {
	// TODO Aufgabe 6: Lade den Zutatennamen und die zugehörigen Cocktails.
	// SELECT name FROM ingredients WHERE ingredient_id = ?
	// SELECT cocktail_id, name FROM cocktails
	// WHERE cocktail_id IN (SELECT cocktail_id FROM instructions WHERE ingredient_id = ?)
	// ORDER BY cocktail_id
	panic("Aufgabe 6")
}

// namePattern erzeugt aus query ein SQL-LIKE-Muster für eine wörtliche Teilzeichenfolge.
func namePattern(query string) string {
	return "%" + strings.NewReplacer("\\", "\\\\", "%", "\\%", "_", "\\_").Replace(query) + "%"
}

// FindCocktailsByName liefert die Cocktails, deren Name den Suchbegriff query enthält.
func (r *DatabaseRepository) FindCocktailsByName(query string) ([]*model.Cocktail, error) {
	return r.listCocktails(`SELECT cocktail_id, name FROM cocktails WHERE name LIKE ? ESCAPE '\' ORDER BY cocktail_id`, namePattern(query))
}

// FindIngredientsByName liefert die Zutaten, deren Name den Suchbegriff query enthält.
func (r *DatabaseRepository) FindIngredientsByName(query string) ([]*model.Ingredient, error) {
	return r.listIngredients(`SELECT ingredient_id, name FROM ingredients WHERE name LIKE ? ESCAPE '\' ORDER BY ingredient_id`, namePattern(query))
}
