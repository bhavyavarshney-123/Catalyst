package extractor

type Extractor struct {
    ai AIService
	prompt string
}

func NewExtractor(ai AIService,prompt string) *Extractor{
	return &Extractor{
		ai: ai
		prompt: prompt,
	}
}

func (e *Extractor) ExtractOpportunity(email models.Email) (ExtractOpportunity,error){


}