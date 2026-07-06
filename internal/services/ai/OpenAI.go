package AI

import (
	"context"

	openai "github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/option"
	"github.com/openai/openai-go/v3/responses"
)

type OpenAIService struct {
	client openai.Client
}

func NewOpenAIService(apiKey string) *OpenAIService {
	client := openai.NewClient(
		option.WithAPIKey(apiKey),
	)

	return &OpenAIService{
		client: client,
	}

}

func (o *OpenAIService) Generate(prompt string, content string) (string, error) {
	ctx := context.Background()

	resp, err := o.client.Responses.New(ctx, responses.ResponseNewParams{
		Input: responses.ResponseNewParamsInputUnion{OfString: openai.String(content)},
		Model: openai.ChatModelGPT5_2,
	})

	if err != nil {
		return " ", err
	}

}
