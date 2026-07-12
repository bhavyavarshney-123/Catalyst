package services

import (
	"context"
	"errors"

	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

type OpenAIEmbeddingService struct {
	client openai.Client
}

func NewOpenAIEmbeddingService(apiKey string) *OpenAIEmbeddingService {
	client := openai.NewClient(
		option.WithAPIKey(apiKey),
	)

	return &OpenAIEmbeddingService{
		client: client,
	}

}

func (o *OpenAIEmbeddingService) Generate(text string) ([]float64, error) {
	ctx := context.Background()

	resp, err := o.client.Embeddings.New(ctx, openai.EmbeddingNewParams{
		Input: openai.EmbeddingNewParamsInputUnion{OfString: openai.String(text)},
		Model: openai.EmbeddingModelTextEmbedding3Small,
	})

	if err != nil {
		return nil, err
	}

	if len(resp.Data) == 0 {
		return nil, errors.New("no embedding returned")
	}

	return resp.Data[0].Embedding, nil
}
