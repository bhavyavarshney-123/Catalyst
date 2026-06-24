package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/bhavyavarshney-123/catalyst/internal/models"
	"github.com/bhavyavarshney-123/catalyst/internal/repository"
	"github.com/go-chi/chi"
)

func GetOpportunity(repo *repository.OpportunityRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		opportunities, err := repo.GetOpportunities()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return

		}

		err = json.NewEncoder(w).Encode(opportunities)
		if err != nil {
			http.Error(w, fmt.Sprintf("Encode error: %v", err), http.StatusBadRequest)
			return
		}
	}
}

func GetOpportunitybyID(repo *repository.OpportunityRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		Opportunity, err := repo.GetOpportunityByID(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		err = json.NewEncoder(w).Encode(Opportunity)
		if err != nil {
			http.Error(w, fmt.Sprintf("Encode error: %v", err), http.StatusBadRequest)
			return
		}
	}
}

func PostOpportunity(repo *repository.OpportunityRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var opportunity models.Opportunity

		err := json.NewDecoder(r.Body).Decode(&opportunity)
		if err != nil {
			http.Error(w, fmt.Sprintf("Decode error: %v", err), http.StatusBadRequest)
			return
		}

		err = repo.CreateOpportunity(opportunity)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write([]byte("Opportunity Created"))

	}
}

func UpdateOpportunity(repo *repository.OpportunityRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var opportunity models.Opportunity

		idStr := chi.URLParam(r, "id")
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = json.NewDecoder(r.Body).Decode(&opportunity)
		if err != nil {
			http.Error(w, fmt.Sprintf("Decode error: %v", err), http.StatusBadRequest)
			return
		}

		err = repo.UpdateOpportunity(opportunity, id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		w.Write([]byte("Opportunity Updated"))

	}
}

func DeleteOpportunity(repo *repository.OpportunityRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		idStr := chi.URLParam(r, "id")
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = repo.DeleteOpportunity(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		w.Write([]byte("Opportunity Deleted"))

	}
}
