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

// NewCocktailHandler erstellt einen Handler. service liefert die Cocktail- und Zutatendaten.
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

// RegisterRoutes registriert die Katalog- und Suchrouten beim übergebenen router.
func (h *CocktailHandler) RegisterRoutes(router *http.ServeMux) {
	// TODO Aufgabe 8: Registriere die Katalogrouten.
	panic("Aufgabe 8")
}

// listCocktails schreibt die Cocktail-Liste als JSON an w. Die Anfrage benötigt keine Parameter.
func (h *CocktailHandler) listCocktails(w http.ResponseWriter, _ *http.Request) {
	// TODO Aufgabe 8: Liefere die Cocktail-Liste.
	panic("Aufgabe 8")
}

// getCocktail liest aus r die Cocktail-ID im Pfad und schreibt das Rezept an w.
func (h *CocktailHandler) getCocktail(w http.ResponseWriter, r *http.Request) {
	// TODO Aufgabe 8: Liefere das angefragte Rezept.
	panic("Aufgabe 8")
}

// listIngredients schreibt die Zutatenliste als JSON an w. Die Anfrage benötigt keine Parameter.
func (h *CocktailHandler) listIngredients(w http.ResponseWriter, _ *http.Request) {
	// TODO Aufgabe 8: Liefere die Zutatenliste.
	panic("Aufgabe 8")
}

// getIngredient liest aus r die Zutaten-ID im Pfad und schreibt den Namen und die passenden Cocktails an w.
func (h *CocktailHandler) getIngredient(w http.ResponseWriter, r *http.Request) {
	// TODO Aufgabe 8: Liefere die angefragte Zutat und ihre Cocktails.
	panic("Aufgabe 8")
}

// search liest den Query-Parameter query aus r und schreibt die Suchtreffer an w.
func (h *CocktailHandler) search(w http.ResponseWriter, r *http.Request) {
	// TODO Aufgabe 8: Liefere die Suchtreffer.
	panic("Aufgabe 8")
}

// resourceID liest eine positive ID aus dem Pfadparameter id von r.
// Bei ungültiger ID schreibt es HTTP 400 an w und gibt false zurück.
func resourceID(w http.ResponseWriter, r *http.Request) (uint, bool) {
	id, err := strconv.ParseUint(r.PathValue("id"), 10, strconv.IntSize)
	if err != nil || id == 0 {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid ID"})
		return 0, false
	}
	return uint(id), true
}

// respond schreibt body als JSON an w. err bestimmt bei einem Fehler den HTTP-Status und den Fehlertext.
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

// writeJSON schreibt body als JSON an w und setzt dabei den HTTP-Statuscode status.
func writeJSON(w http.ResponseWriter, status int, body any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	// TODO Aufgabe 8: Schreibe das Antwortobjekt als JSON.
}
