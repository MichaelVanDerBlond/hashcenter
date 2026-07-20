package workflow

import "time"

type Context struct {
	ID string

	Capture string

	HashFile string

	Dictionaries []string

	OutputDir string

	StartedAt time.Time

	FinishedAt time.Time

	CurrentStage string

	CompletedStages []string

	Values map[string]any
}

func NewContext() *Context {
	return &Context{
		Values: make(map[string]any),
	}
}
