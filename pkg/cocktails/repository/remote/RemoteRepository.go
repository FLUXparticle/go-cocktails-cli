package remote

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"go-cocktails-cli/pkg/cocktails/model"
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
	req, err := http.NewRequestWithContext(r.ctx, method, r.baseURL+path, bytes.NewReader(body))
	if err != nil {
		return err
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	resp, err := r.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("%s %s: %s", method, path, resp.Status)
	}
	if result == nil {
		return nil
	}
	return json.NewDecoder(resp.Body).Decode(result)
}

func (r *RemoteRepository) GetCocktails() ([]*model.Cocktail, error) {
	var rows []*model.Cocktail
	err := r.request(http.MethodGet, "/api/cocktails", nil, &rows)
	return rows, err
}

func (r *RemoteRepository) GetCocktail(id uint) (*model.Cocktail, error) {
	var row model.Cocktail
	err := r.request(http.MethodGet, fmt.Sprintf("/api/cocktails/%d", id), nil, &row)
	return &row, err
}

func (r *RemoteRepository) GetIngredients() ([]*model.Ingredient, error) {
	var rows []*model.Ingredient
	err := r.request(http.MethodGet, "/api/ingredients", nil, &rows)
	return rows, err
}

func (r *RemoteRepository) GetCocktailsByIngredient(id uint) (string, []*model.Cocktail, error) {
	var result struct {
		Name      string            `json:"name"`
		Cocktails []*model.Cocktail `json:"cocktails"`
	}
	err := r.request(http.MethodGet, fmt.Sprintf("/api/ingredients/%d", id), nil, &result)
	return result.Name, result.Cocktails, err
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
	rows, err := r.getFridge()
	if err != nil {
		return nil, err
	}
	stock := make(map[uint]bool)
	for _, row := range rows {
		if row.InFridge {
			stock[row.IngredientID] = true
		}
	}
	return stock, nil
}

func (r *RemoteRepository) getFridge() ([]*model.FridgeIngredient, error) {
	var rows []*model.FridgeIngredient
	err := r.request(http.MethodGet, "/api/user/"+url.PathEscape(fridgeUser)+"/fridge", nil, &rows)
	return rows, err
}

func (r *RemoteRepository) SetFridgeIngredient(id uint, inFridge bool) error {
	path := fmt.Sprintf("/api/user/%s/fridge/%d", url.PathEscape(fridgeUser), id)
	return r.request(http.MethodPatch, path, []byte(fmt.Sprintf(`{"inFridge":%t}`, inFridge)), nil)
}
