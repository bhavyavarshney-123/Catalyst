package main

import (
	"fmt"
	"net/http"

	"github.com/bhavyavarshney-123/catalyst/internal/database"
	"github.com/bhavyavarshney-123/catalyst/internal/handlers"
	"github.com/bhavyavarshney-123/catalyst/internal/repository"
	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)

func main() {
	r := chi.NewRouter()

	err := godotenv.Load()

	if err != nil {

		fmt.Println("Error Loading the .env")
	}

	db, err := database.Connect()
	if err != nil {
		fmt.Println(err)
		return
	}

	defer db.Close()

	repo := repository.NewOpportunityRepository(db)

	r.Post("/opportunities", handlers.PostOpportunity(repo))
	r.Get("/opportunities", handlers.GetOpportunity(repo))
	r.Get("/opportunities/{id}", handlers.GetOpportunitybyID(repo))
	r.Put("/opportunities/{id}", handlers.UpdateOpportunity(repo))
	r.Delete("/opportunities/{id}", handlers.DeleteOpportunity(repo))
	fmt.Println("Server started on :8080")

	err = http.ListenAndServe(":8080", r)
	if err != nil {
		fmt.Println(err)
	}

}
