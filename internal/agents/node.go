package agents

import "context"

type Node interface {
	Name() string
	Execute(ctx context.Context, state *State) error
}
