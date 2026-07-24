package agents

type ConditionFunc func(*State) bool

type Edge struct {
	From      string
	To        string
	Condition ConditionFunc
}
