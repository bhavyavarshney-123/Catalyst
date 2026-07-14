package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/bhavyavarshney-123/catalyst/internal/services/rag"
)

func CLI(ragService *rag.RAGService) {
	fmt.Println("=========================")
	fmt.Println("Welcome to Catalyst")
	fmt.Println("Ask me anything about your job opportunities.")
	fmt.Println("Type 'exit' to quit.")
	fmt.Println("=========================")

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("> ")

		if !scanner.Scan() {
			if err := scanner.Err(); err != nil {
				fmt.Println("Error:", err)
			}
			break
		}

		question := strings.TrimSpace(scanner.Text())

		if question == "" {
			continue
		}

		if strings.EqualFold(question, "exit") {
			fmt.Println("Goodbye!")
			return
		}

		answer, err := ragService.Answer(question)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		fmt.Println()
		fmt.Println(answer)
		fmt.Println()
	}
}
