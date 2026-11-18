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

// run zeigt die Cocktail-Liste oder ein Rezept an. Eine id von 0 wählt die Liste.
func run(id uint) error {
	// TODO Aufgabe 4: Zeige Cocktails vom Kursserver an.
	panic("Aufgabe 4")
}
