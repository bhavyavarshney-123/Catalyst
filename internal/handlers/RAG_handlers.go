package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/bhavyavarshney-123/catalyst/internal/services/rag"
)

type QueryRequest struct {
	Question string `json:"question"`
}

type QueryResponse struct {
	Answer string `json:"answer"`
}

func RAG(RAGService *rag.RAGService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var req QueryRequest

		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, "invalid request body", http.StatusBadRequest)
			return
		}

		if strings.TrimSpace(req.Question) == "" {
			http.Error(w, "question is required", http.StatusBadRequest)
			return
		}

		answer, err := RAGService.Answer(req.Question)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
		}
		response := QueryResponse{
			Answer: answer,
		}
		err = json.NewEncoder(w).Encode(response)
		if err != nil {
			http.Error(w, fmt.Sprintf("Encode error: %v", err), http.StatusBadRequest)
			return
		}

	}
}
