package handler

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"go-cocktails-cli/pkg/cocktails/model"
	"go-cocktails-cli/pkg/cocktails/repository"
	"go-cocktails-cli/pkg/cocktails/service"
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
	router.HandleFunc("GET /api/cocktails", h.listCocktails)
	router.HandleFunc("GET /api/cocktails/{id}", h.getCocktail)
	router.HandleFunc("GET /api/ingredients", h.listIngredients)
	router.HandleFunc("GET /api/ingredients/{id}", h.getIngredient)
	router.HandleFunc("GET /api/search", h.search)
}

func (h *CocktailHandler) listCocktails(w http.ResponseWriter, _ *http.Request) {
	rows, err := h.service.GetCocktails()
	respond(w, rows, err)
}

func (h *CocktailHandler) getCocktail(w http.ResponseWriter, r *http.Request) {
	id, ok := resourceID(w, r)
	if !ok {
		return
	}
	row, err := h.service.GetCocktail(id)
	if err != nil {
		respond(w, nil, err)
		return
	}
	body := cocktailJSON{ID: row.CocktailID, Name: row.Name, Instructions: make([]instructionJSON, 0, len(row.Instructions))}
	for _, instruction := range row.Instructions {
		body.Instructions = append(body.Instructions, instructionJSON{Amount: instruction.AmountCL, IngredientID: instruction.IngredientID, Ingredient: instruction.Ingredient})
	}
	respond(w, body, nil)
}

func (h *CocktailHandler) listIngredients(w http.ResponseWriter, _ *http.Request) {
	rows, err := h.service.GetIngredients()
	respond(w, rows, err)
}

func (h *CocktailHandler) getIngredient(w http.ResponseWriter, r *http.Request) {
	id, ok := resourceID(w, r)
	if !ok {
		return
	}
	name, cocktails, err := h.service.GetCocktailsByIngredient(id)
	respond(w, map[string]any{"name": name, "cocktails": cocktails}, err)
}

func (h *CocktailHandler) search(w http.ResponseWriter, r *http.Request) {
	rows, err := h.service.Search(r.URL.Query().Get("query"))
	respond(w, rows, err)
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
	if err := json.NewEncoder(w).Encode(body); err != nil {
		log.Print(err)
	}
}
