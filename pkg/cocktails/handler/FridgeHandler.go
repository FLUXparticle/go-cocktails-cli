package handler

import (
	"go-cocktails-cli/pkg/cocktails/service"
	"net/http"
)

type FridgeHandler struct {
	fridge *service.FridgeService
}

// NewFridgeHandler erstellt einen Handler. fridge liest und ändert den Kühlschrankbestand.
func NewFridgeHandler(fridge *service.FridgeService) *FridgeHandler {
	return &FridgeHandler{fridge: fridge}
}

// RegisterRoutes registriert die Kühlschrankrouten beim übergebenen router.
func (h *FridgeHandler) RegisterRoutes(router *http.ServeMux) {
	// TODO Aufgabe 8: Registriere die Kühlschrankrouten für den Bonus.
}

// list schreibt die Zutaten samt Kühlschrankstatus an w. Der Bestand ist für alle Benutzer gemeinsam.
func (h *FridgeHandler) list(w http.ResponseWriter, _ *http.Request) {
	// TODO Aufgabe 8: Liefere den Kühlschrankbestand.
	panic("Aufgabe 8")
}

// update liest aus r die Zutaten-ID im Pfad und den Bool-Wert inFridge aus dem JSON-Body.
// Die Methode schreibt das Ergebnis an w.
func (h *FridgeHandler) update(w http.ResponseWriter, r *http.Request) {
	// TODO Aufgabe 8: Ändere eine Zutat im Kühlschrank.
	panic("Aufgabe 8")
}

// possible schreibt die mit dem aktuellen Kühlschrankbestand mixbaren Cocktails an w.
func (h *FridgeHandler) possible(w http.ResponseWriter, _ *http.Request) {
	// TODO Aufgabe 8: Liefere die möglichen Cocktails als HTTP-Antwort.
	panic("Aufgabe 8")
}
