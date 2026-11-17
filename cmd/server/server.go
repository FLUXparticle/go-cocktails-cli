package main

import (
	"fmt"
	"net"
	"net/http"
	"os"

	"go-cocktails-cli/pkg/cocktails/handler"
	"go-cocktails-cli/pkg/cocktails/repository/database"
	"go-cocktails-cli/pkg/cocktails/repository/memory"
	"go-cocktails-cli/pkg/cocktails/service"
	"go-cocktails-cli/web"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	repo, err := database.NewDatabaseRepository()
	if err != nil {
		return err
	}
	defer repo.Close()

	router := newRouter(repo)
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		return err
	}
	fmt.Println("Server läuft unter http://localhost:8080/")
	return http.Serve(listener, router)
}

func newRouter(repo *database.DatabaseRepository) *http.ServeMux {
	router := http.NewServeMux()
	serv := service.NewCocktailService(repo)
	handler.NewCocktailHandler(serv).RegisterRoutes(router)
	fridge := service.NewFridgeService(memory.NewFridgeRepository(), serv)
	handler.NewFridgeHandler(fridge).RegisterRoutes(router)
	router.Handle("GET /", http.FileServerFS(web.Files))
	return router
}
