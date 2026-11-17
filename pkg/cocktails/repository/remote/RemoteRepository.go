package remote

import (
	"context"
	"errors"
	"fmt"
	"go-cocktails-cli/pkg/cocktails/model"
	"net/http"
	"strings"
	"time"
)

const (
	fridgeUser = ""
	courseURL  = "https://cocktails.fluxparticle.com"
)

type RemoteRepository struct {
	ctx     context.Context
	baseURL string
	client  *http.Client
}

func NewRemoteRepository() *RemoteRepository {
	return NewRemoteRepositoryWithContext(context.Background())
}

func NewRemoteRepositoryWithContext(ctx context.Context) *RemoteRepository {
	return NewRemoteRepositoryWithClient(ctx, &http.Client{Timeout: 15 * time.Second})
}

func NewRemoteRepositoryWithClient(ctx context.Context, client *http.Client) *RemoteRepository {
	return NewRemoteRepositoryWithURL(ctx, client, courseURL)
}

func NewRemoteRepositoryWithURL(ctx context.Context, client *http.Client, baseURL string) *RemoteRepository {
	if strings.TrimSpace(fridgeUser) == "" {
		panic("Setze die Konstante fridgeUser in RemoteRepository.go auf deine Benutzer-ID")
	}
	return &RemoteRepository{ctx: ctx, baseURL: baseURL, client: client}
}

func (r *RemoteRepository) request(method string, path string, body []byte, result any) error {
	// TODO Aufgabe 1: Führe den HTTP-Request aus und dekodiere die Antwort.
	panic("Aufgabe 1")
}

func (r *RemoteRepository) GetCocktails() ([]*model.Cocktail, error) {
	// TODO Aufgabe 1: Lade die Cocktail-Liste vom Kursserver.
	panic("Aufgabe 1")
}

func (r *RemoteRepository) GetCocktail(id uint) (*model.Cocktail, error) {
	// TODO Aufgabe 1: Lade das Rezept für eine Cocktail-ID.
	panic("Aufgabe 1")
}

func (r *RemoteRepository) GetIngredients() ([]*model.Ingredient, error) {
	// TODO Aufgabe 4: Lade die Zutatenliste vom Kursserver.
	panic("Aufgabe 4")
}

func (r *RemoteRepository) GetCocktailsByIngredient(id uint) (string, []*model.Cocktail, error) {
	// TODO Aufgabe 4: Lade die Cocktails zu einer Zutaten-ID.
	panic("Aufgabe 4")
}

// FindCocktailsByName is unsupported by the remote repository.
func (r *RemoteRepository) FindCocktailsByName(query string) ([]*model.Cocktail, error) {
	return nil, fmt.Errorf("RemoteRepository.FindCocktailsByName(%q): %w", query, errors.ErrUnsupported)
}

// FindIngredientsByName is unsupported by the remote repository.
func (r *RemoteRepository) FindIngredientsByName(query string) ([]*model.Ingredient, error) {
	return nil, fmt.Errorf("RemoteRepository.FindIngredientsByName(%q): %w", query, errors.ErrUnsupported)
}

func (r *RemoteRepository) Ingredients() (map[uint]bool, error) {
	// TODO Aufgabe 5: Lies die vorhandenen Zutaten.
	panic("Aufgabe 5")
}

func (r *RemoteRepository) getFridge() ([]*model.FridgeIngredient, error) {
	// TODO Aufgabe 5: Lade den Kühlschrankbestand.
	panic("Aufgabe 5")
}

func (r *RemoteRepository) SetFridgeIngredient(id uint, inFridge bool) error {
	// TODO Aufgabe 5: Ändere den Bestand beim Kursserver.
	panic("Aufgabe 5")
}
