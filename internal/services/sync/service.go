package sync

import (
	"errors"

	"github.com/bhavyavarshney-123/catalyst/internal/extractor"
	"github.com/bhavyavarshney-123/catalyst/internal/repository"
	"github.com/bhavyavarshney-123/catalyst/internal/services/gmail"
)

type SyncService struct {
	gmail     *gmail.GmailManager
	extractor *extractor.Extractor
	repo      *repository.OpportunityRepository
}

func NewSyncService(gmail *gmail.GmailManager, extractor *extractor.Extractor, repo *repository.OpportunityRepository) *SyncService {
	return &SyncService{
		gmail:     gmail,
		extractor: extractor,
		repo:      repo,
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
		opportunity, err := extractor.ToOpportunity(extracted)
		if err != nil {
			return err
		}
		if err := s.repo.CreateOpportunity(opportunity); err != nil {
			return err
		}
	}

	return nil
}
