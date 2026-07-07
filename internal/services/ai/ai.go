package ai

type AIService interface {
	Generate(prompt string, content string) (string, error)
}
