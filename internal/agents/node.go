package agent

import "context"

func SyncGmail(ctx context.Context, state State) (State, error)

func ExtractOpportunities(ctx context.Context, state State) (State, error)

func FindRelevantOpportunities(ctx context.Context, state State) (State, error)

func GenerateResponse(ctx context.Context, state State) (State, error)
