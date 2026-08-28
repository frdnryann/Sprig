package main

import (
	"log"
	"net/http"
	"sprig/internal/config"
	"sprig/internal/features/categories"
	"sprig/internal/features/users"
)

func main() {
	cfg := config.LoadConfig()

	// mysql connection
	db, err := config.OpenConnection(cfg)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	mux := http.NewServeMux()

	// Dependency Injection
	userRepository := users.NewUserRepository(db)
	userService := users.NewUserService(userRepository)
	userHandler := users.NewUserHandler(userService)

	// ENDPOINT
	mux.HandleFunc("POST /api/users", userHandler.Create)
	mux.HandleFunc("GET /api/users/{id}", userHandler.FindByID)
	mux.HandleFunc("GET /api/users", userHandler.FindAll)
	mux.HandleFunc("PATCH /api/users/id", userHandler.Update)
	mux.HandleFunc("DELETE /api/users/{id}", userHandler.Delete)

	server := http.Server{
		Addr:    ":8080", // containerization (cukup ambil portnya saja)
		Handler: mux,
	}

	log.Printf("Server berjalan di port %s", server.Addr)
	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
