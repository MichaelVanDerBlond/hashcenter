package workflow

import (
	"context"
	"errors"

	"github.com/MichaelVanDerBlond/hashcenter/internal/analyze"
)

type AnalyzeStage struct{}

func NewAnalyzeStage() *AnalyzeStage {
	return &AnalyzeStage{}
}

func (a *AnalyzeStage) Name() string {
	return "analyze"
}

func (a *AnalyzeStage) ShouldRun(*Context) bool {
	return true
}

func (a *AnalyzeStage) Run(ctx context.Context, wf *Context) error {
	_ = ctx

	if wf == nil {
		return errors.New("workflow context is nil")
	}

	if wf.HashFile == "" {
		return errors.New("hash file is not specified")
	}

	report, err := analyze.Analyze(wf.HashFile)
	if err != nil {
		return err
	}

	wf.Analyze = report

	return nil
}
