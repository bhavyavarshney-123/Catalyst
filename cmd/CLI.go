package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	embeddings "github.com/bhavyavarshney-123/catalyst/internal/services/embeddings"
)

func CLI(embeddings embeddings.EmbeddingService) {
	fmt.Println("=========================")
	fmt.Println("Welcome to Catalyst")
	fmt.Println("=========================")

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("> ")

		if !scanner.Scan() {
			if err := scanner.Err(); err != nil {
				fmt.Println(err)
			}
			break
		}

		question := scanner.Text()

		question = strings.TrimSpace(question)

		if question == "" {
			continue
		}

		if question == "exit" {
			break
		}

		embedding, err := embeddings.Generate(question)

		if err != nil {
			fmt.Println("failed to generate embedding:", err)
			continue

		}

	}

	fmt.Println("Goodbye!")
}
