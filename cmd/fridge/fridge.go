package main

import (
	"fmt"
	"os"
)

func main() {
	// TODO Aufgabe 5: Lies -add, -remove, -display und -possible ein.
	if err := run(0, 0, false, false); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

// run führt die gewählten Kühlschrankaktionen der Reihe nach aus.
// add und remove sind Zutaten-IDs. Der Wert 0 überspringt die jeweilige Aktion.
// display zeigt den Bestand an. possible zeigt die daraus mixbaren Cocktails.
func run(add uint, remove uint, display bool, possible bool) error {
	// TODO Aufgabe 5: Führe die gewählten Kühlschrankaktionen aus.
	panic("Aufgabe 5")
}
