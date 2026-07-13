package rag

import (
	"github.com/bhavyavarshney-123/catalyst/internal/repository"
	"github.com/bhavyavarshney-123/catalyst/internal/services/ai"
	embeddings "github.com/bhavyavarshney-123/catalyst/internal/services/embeddings"
)

type RAGService struct {
	ai        ai.AIService
	repo      *repository.OpportunityRepository
	embedding embeddings.EmbeddingService
}

func NewRaGService(ai ai.AIService, repo *repository.OpportunityRepository, embedding embeddings.EmbeddingService) *RAGService {
	return &RAGService{ai: ai, repo: repo, embedding: embedding}
}

func (r *RAGService) Answer(question string) (string, error) {
	embedding, err := r.embedding.Generate(question)

	if err != nil {
		return "", err
	}

	results, err := r.repo.SearchSimilar(embedding, 5)

	context := BuildContext(results)

	RAGUserPrompt := BuildRAGUserPrompt(question, context)

	answer, err := r.ai.Generate(RAGSystemPrompt, RAGUserPrompt)
	if err != nil {
		return "", err
	}

	return answer, nil

}
