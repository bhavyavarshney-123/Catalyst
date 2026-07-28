package cli

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/bhavyavarshney-123/catalyst/internal/agents"
)

func CLI(agent *agents.Agent) {
	fmt.Println("=========================")
	fmt.Println("Welcome to Catalyst")
	fmt.Println("Ask me anything about your job opportunities.")
	fmt.Println("Type 'exit' to quit.")
	fmt.Println("=========================")

	graph := agents.BuildGraph(agent)
	ctx := context.Background()

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

		state := &agents.State{
			UserQuestion: question,
			Limit:        2,
		}

		if err := graph.Execute(ctx, state); err != nil {
			fmt.Println("Error:", err)
			continue
		}

		fmt.Println(state.Response)
	}
}
