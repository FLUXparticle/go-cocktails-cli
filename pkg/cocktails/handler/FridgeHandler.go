package handler

import (
	"encoding/json"
	"net/http"

	"go-cocktails-cli/pkg/cocktails/service"
)

type FridgeHandler struct {
	fridge *service.FridgeService
}

func NewFridgeHandler(fridge *service.FridgeService) *FridgeHandler {
	return &FridgeHandler{fridge: fridge}
}

func (h *FridgeHandler) RegisterRoutes(router *http.ServeMux) {
	// The client supplies a user ID, but this server has one shared fridge.
	router.HandleFunc("GET /api/user/{user}/fridge", h.list)
	router.HandleFunc("PATCH /api/user/{user}/fridge/{id}", h.update)
	router.HandleFunc("GET /api/user/{user}/possible", h.possible)
}

func (h *FridgeHandler) list(w http.ResponseWriter, _ *http.Request) {
	rows, err := h.fridge.GetFridge()
	respond(w, rows, err)
}

func (h *FridgeHandler) update(w http.ResponseWriter, r *http.Request) {
	id, ok := resourceID(w, r)
	if !ok {
		return
	}
	var body struct {
		InFridge *bool `json:"inFridge"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body.InFridge == nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "expected inFridge boolean"})
		return
	}
	respond(w, nil, h.fridge.SetFridgeIngredient(id, *body.InFridge))
}

func (h *FridgeHandler) possible(w http.ResponseWriter, _ *http.Request) {
	rows, err := h.fridge.GetPossibleCocktails()
	respond(w, rows, err)
}
