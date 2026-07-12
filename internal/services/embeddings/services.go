package services

type EmbeddingService interface {
	Generate(text string) ([]float64, error)
}
