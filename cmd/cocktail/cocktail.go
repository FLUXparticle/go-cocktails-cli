package main

import (
	"fmt"
	"os"
)

func main() {
	// TODO Aufgabe 4: Lies die Cocktail-ID mit -id oder -i ein.
	if err := run(0); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(id uint) error {
	// TODO Aufgabe 4: Zeige Cocktails vom Kursserver an.
	panic("Aufgabe 4")
}
