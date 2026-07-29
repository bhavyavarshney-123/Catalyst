package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	OpenAIKey        string
	DatabaseURL      string
	GmailCredentials string
}

func Load() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("loading .env: %w", err)
	}

	cfg := &Config{
		OpenAIKey:        os.Getenv("OPENAI_API_KEY"),
		DatabaseURL:      os.Getenv("DATABASE_URL"),
		GmailCredentials: os.Getenv("GMAIL_CREDENTIALS_FILE"),
	}

	if cfg.OpenAIKey == "" {
		return nil, fmt.Errorf("OPENAI_API_KEY is required")
	}

	if cfg.DatabaseURL == "" {
		return nil, fmt.Errorf("DATABASE_URL is required")
	}

	if cfg.GmailCredentials == "" {
		return nil, fmt.Errorf("GMAIL_CREDENTIALS_FILE is required")
	}

	return cfg, nil
}
