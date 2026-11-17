package handler

import (
	"errors"
	"go-cocktails-cli/pkg/cocktails/model"
	"go-cocktails-cli/pkg/cocktails/repository"
	"go-cocktails-cli/pkg/cocktails/service"
	"log"
	"net/http"
	"strconv"
)

type CocktailHandler struct {
	service *service.CocktailService
}

func NewCocktailHandler(service *service.CocktailService) *CocktailHandler {
	return &CocktailHandler{service: service}
}

type instructionJSON struct {
	Ingredient   *model.Ingredient `json:"ingredient"`
	Amount       int               `json:"amount"`
	IngredientID uint              `json:"ingredient_id"`
}

type cocktailJSON struct {
	ID           uint              `json:"id,omitempty"`
	Name         string            `json:"name"`
	Instructions []instructionJSON `json:"instructions"`
}

func (h *CocktailHandler) RegisterRoutes(router *http.ServeMux) {
	// TODO Aufgabe 8: Registriere die Katalogrouten.
	panic("Aufgabe 8")
}

func (h *CocktailHandler) listCocktails(w http.ResponseWriter, _ *http.Request) {
	// TODO Aufgabe 8: Liefere die Cocktail-Liste.
	panic("Aufgabe 8")
}

func (h *CocktailHandler) getCocktail(w http.ResponseWriter, r *http.Request) {
	// TODO Aufgabe 8: Liefere das angefragte Rezept.
	panic("Aufgabe 8")
}

func (h *CocktailHandler) listIngredients(w http.ResponseWriter, _ *http.Request) {
	// TODO Aufgabe 8: Liefere die Zutatenliste.
	panic("Aufgabe 8")
}

func (h *CocktailHandler) getIngredient(w http.ResponseWriter, r *http.Request) {
	// TODO Aufgabe 8: Liefere die angefragte Zutat und ihre Cocktails.
	panic("Aufgabe 8")
}

func (h *CocktailHandler) search(w http.ResponseWriter, r *http.Request) {
	// TODO Aufgabe 8: Liefere die Suchtreffer.
	panic("Aufgabe 8")
}

func resourceID(w http.ResponseWriter, r *http.Request) (uint, bool) {
	id, err := strconv.ParseUint(r.PathValue("id"), 10, strconv.IntSize)
	if err != nil || id == 0 {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid ID"})
		return 0, false
	}
	return uint(id), true
}

func respond(w http.ResponseWriter, body any, err error) {
	switch {
	case errors.Is(err, repository.ErrNotFound):
		writeJSON(w, http.StatusNotFound, map[string]string{"error": err.Error()})
	case err != nil:
		log.Print(err)
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": "internal server error"})
	default:
		writeJSON(w, http.StatusOK, body)
	}
}

func writeJSON(w http.ResponseWriter, status int, body any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	// TODO Aufgabe 8: Schreibe das Antwortobjekt als JSON.
}
