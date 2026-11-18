package service

import (
	"testing"
)

// TestMilkAmount prüft die Milchmenge beim sequenziellen Abruf gegen den Kursserver.
func TestMilkAmount(t *testing.T) {
	// TODO Aufgabe 1: Prüfe die Milchmenge gegen den Kursserver.
	t.Skip("Aufgabe 1")
}

// TestMilkAmountParallel prüft die Milchmenge beim parallelen Abruf gegen den Kursserver.
func TestMilkAmountParallel(t *testing.T) {
	// TODO Aufgabe 2: Prüfe die parallel ermittelte Milchmenge.
	t.Skip("Aufgabe 2")
}

// TestMilkAmountErrGroupCancelsRequests prüft, dass der Fehler für die unbekannte Cocktail-ID vom Kursserver zurückgegeben wird.
func TestMilkAmountErrGroupCancelsRequests(t *testing.T) {
	// TODO Aufgabe 3: Prüfe den Fehler beim Abruf gegen den Kursserver.
	t.Skip("Aufgabe 3")
}
