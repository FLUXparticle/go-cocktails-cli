package service

import (
	"go-cocktails-cli/pkg/cocktails/repository/remote"
	"testing"
)

func TestMilkAmount(t *testing.T) {
	repo := remote.NewRemoteRepository()
	svc := NewCocktailService(repo)

	got, err := svc.MilkAmount()
	if err != nil {
		t.Fatal(err)
	}
	const want = 38
	if got != want {
		t.Errorf("Milchmenge: %d cl, erwartet: %d cl", got, want)
	}
}

func TestMilkAmountParallel(t *testing.T) {
	repo := remote.NewRemoteRepository()
	svc := NewCocktailService(repo)

	got, err := svc.MilkAmountParallel()
	if err != nil {
		t.Fatal(err)
	}
	const want = 38
	if got != want {
		t.Errorf("Milchmenge: %d cl, erwartet: %d cl", got, want)
	}
}

func TestMilkAmountErrGroupCancelsRequests(t *testing.T) {
	// TODO Aufgabe 3: Prüfe den Fehler beim Abruf gegen den Kursserver.
	t.Skip("Aufgabe 3")
}
