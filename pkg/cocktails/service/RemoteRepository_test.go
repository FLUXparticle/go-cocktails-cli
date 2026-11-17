package service

import (
	"net/http"
	"strings"
	"testing"

	"go-cocktails-cli/pkg/cocktails/repository/remote"
	"golang.org/x/sync/errgroup"
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
	group, ctx := errgroup.WithContext(t.Context())
	client := &http.Client{}
	repo := remote.NewRemoteRepositoryWithClient(ctx, client)
	svc := NewCocktailService(repo)

	_, err := svc.MilkAmountErrGroup(group)
	if err == nil || !strings.Contains(err.Error(), "/api/cocktails/70: 404") {
		t.Fatalf("expected HTTP 404 for cocktail 70, got %v", err)
	}
}
