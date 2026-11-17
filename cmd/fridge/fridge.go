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

func run(add uint, remove uint, display bool, possible bool) error {
	// TODO Aufgabe 5: Führe die gewählten Kühlschrankaktionen aus.
	panic("Aufgabe 5")
}
