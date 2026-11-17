package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"go-cocktails-cli/pkg/cocktails/model"
	"go-cocktails-cli/pkg/cocktails/repository/database"
	"go-cocktails-cli/pkg/cocktails/service"
)

func TestCocktailList(t *testing.T) {
	repo, err := database.NewDatabaseRepository()
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = repo.Close() })

	router := http.NewServeMux()
	NewCocktailHandler(service.NewCocktailService(repo)).RegisterRoutes(router)
	server := httptest.NewTestServer(t, router)
	client := server.Client()
	response, err := client.Get(server.URL + "/api/cocktails")
	if err != nil {
		t.Fatal(err)
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		t.Fatalf("HTTP-Status: %d, erwartet: %d", response.StatusCode, http.StatusOK)
	}

	var cocktails []*model.Cocktail
	if err := json.NewDecoder(response.Body).Decode(&cocktails); err != nil {
		t.Fatal(err)
	}
	const want = 69
	if got := len(cocktails); got != want {
		t.Errorf("Cocktails: %d, erwartet: %d", got, want)
	}
}
