package ai

type OpenAIService struct {
	client *openai.Client
}

func NewOpenAIService(apiKey string) *OpenAIService {

}

func (o *OpenAIService) Generate(prompt string, content string) (string, error) {

}
