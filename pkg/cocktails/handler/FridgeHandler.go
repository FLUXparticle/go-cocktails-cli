package handler

import (
	"go-cocktails-cli/pkg/cocktails/service"
	"net/http"
)

type FridgeHandler struct {
	fridge *service.FridgeService
}

func NewFridgeHandler(fridge *service.FridgeService) *FridgeHandler {
	return &FridgeHandler{fridge: fridge}
}

func (h *FridgeHandler) RegisterRoutes(router *http.ServeMux) {
	// TODO Aufgabe 8: Registriere die Kühlschrankrouten für den Bonus.
}

func (h *FridgeHandler) list(w http.ResponseWriter, _ *http.Request) {
	// TODO Aufgabe 8: Liefere den Kühlschrankbestand.
	panic("Aufgabe 8")
}

func (h *FridgeHandler) update(w http.ResponseWriter, r *http.Request) {
	// TODO Aufgabe 8: Ändere eine Zutat im Kühlschrank.
	panic("Aufgabe 8")
}

func (h *FridgeHandler) possible(w http.ResponseWriter, _ *http.Request) {
	// TODO Aufgabe 8: Liefere die möglichen Cocktails als HTTP-Antwort.
	panic("Aufgabe 8")
}
