package main

import (
	"fmt"
	"os"
)

func main() {
	// TODO Aufgabe 4: Lies die Zutaten-ID mit -id oder -i ein.
	if err := run(0); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

// run zeigt bei id 0 die Zutatenliste. Sonst zeigt es die Zutat mit dieser ID und die Cocktails, die sie enthalten.
func run(id uint) error {
	// TODO Aufgabe 4: Zeige Zutaten vom Kursserver an.
	panic("Aufgabe 4")
}
