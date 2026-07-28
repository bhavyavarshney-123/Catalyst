package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/bhavyavarshney-123/catalyst/cmd/cli"
	app "github.com/bhavyavarshney-123/catalyst/internal"
)

func main() {
	app, err := app.New()
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		fmt.Println("Server started on :8080")
		log.Fatal(http.ListenAndServe(":8080", app.Router))
	}()

	cli.CLI(app.Agent)
}
