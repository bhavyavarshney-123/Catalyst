package sync

import (
	"errors"
	"fmt"

	"github.com/bhavyavarshney-123/catalyst/internal/extractor"
	"github.com/bhavyavarshney-123/catalyst/internal/repository"
	embeddings "github.com/bhavyavarshney-123/catalyst/internal/services/embeddings"
	"github.com/bhavyavarshney-123/catalyst/internal/services/gmail"
)

type SyncService struct {
	gmail     *gmail.GmailManager
	extractor *extractor.Extractor
	repo      *repository.OpportunityRepository
	embedding embeddings.EmbeddingService
}

func NewSyncService(gmail *gmail.GmailManager, extractor *extractor.Extractor, repo *repository.OpportunityRepository, embedding embeddings.EmbeddingService) *SyncService {
	return &SyncService{
		gmail:     gmail,
		extractor: extractor,
		repo:      repo,
		embedding: embedding,
	}
}

func (s *SyncService) Sync(limit int) error {
	if s.gmail == nil {
		return errors.New("gmail not connected")
	}
	emails, err := s.gmail.Service.ListRecentMessages(int64(limit))
	if err != nil {
		return err
	}

	for _, email := range emails {

		extracted, err := s.extractor.ExtractOpportunity(email)
		if err != nil {
			return err
		}

		if !extracted.IsOpportunity {
			continue
		}

		opportunitytxt := buildEmbeddingText(extracted)

		embedding, err := s.embedding.Generate(opportunitytxt)

		if err != nil {
			return err
		}

		opportunity, err := extractor.ToOpportunity(extracted, embedding)
		if err != nil {
			return err
		}

		id, err := s.repo.CheckExists(opportunity.CompanyName, opportunity.RoleApplied)

		if err != nil {
			return err
		}

		if id == 0 {
			fmt.Println("Opportunity does not exist, creating a new opportunity")

			if err := s.repo.CreateOpportunity(opportunity); err != nil {
				return err
			}
		} else {

			fmt.Println("Opportunity already exists, updating it")

			if err := s.repo.UpdateOpportunity(opportunity, id); err != nil {
				return err
			}

		}

	}

	return nil
}
