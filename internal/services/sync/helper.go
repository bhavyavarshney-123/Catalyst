package sync

import (
	"fmt"

	extractor "github.com/bhavyavarshney-123/catalyst/internal/extractor/models"
)

func buildEmbeddingText(extractedopportunity extractor.ExtractOpportunity) string {
	return fmt.Sprintf(
		`Company: %s
Role: %s
Status: %s
Comments: %s`,
		extractedopportunity.CompanyName,
		extractedopportunity.RoleApplied,
		extractedopportunity.ApplicationStatus,
		extractedopportunity.Comments,
	)
}
