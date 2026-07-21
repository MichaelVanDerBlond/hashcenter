package workflow

import (
	"time"

	"github.com/MichaelVanDerBlond/hashcenter/internal/analyze"
	"github.com/MichaelVanDerBlond/hashcenter/internal/convert"
	"github.com/MichaelVanDerBlond/hashcenter/internal/jobs"
)

type Context struct {
	ID string

	SessionID string

	Capture string

	HashFile string

	Dictionary string

	Dictionaries []string

	AttackMode int

	HashMode int

	Extra []string

	OutputDir string

	StartedAt time.Time

	FinishedAt time.Time

	CurrentStage string

	CompletedStages []string

	Analyze *analyze.Report

	Convert *convert.Result

	Result *jobs.Result

	Values map[string]any
}

func NewContext() *Context {
	return &Context{
		Values: make(map[string]any),
	}
}
