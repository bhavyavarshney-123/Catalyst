package ai

import (
	"context"

	openai "github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
	"github.com/openai/openai-go/responses"
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
		Instructions: openai.String(prompt),
		Input:        responses.ResponseNewParamsInputUnion{OfString: openai.String(content)},
		Model:        openai.ChatModelGPT4o,
	})

	if err != nil {
		return "", err
	}
	textresponse := resp.OutputText()
	return textresponse, nil

}
