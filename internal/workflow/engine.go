package workflow

import (
	"context"
	"fmt"
)

type Engine struct {
	stages []Stage
}

func NewEngine() *Engine {
	return &Engine{
		stages: make([]Stage, 0),
	}
}

func (e *Engine) Add(stage Stage) *Engine {
	e.stages = append(e.stages, stage)
	return e
}

func (e *Engine) Run(ctx context.Context, wf *Context) error {

	for _, stage := range e.stages {

		if !stage.ShouldRun(wf) {
			continue
		}

		wf.CurrentStage = stage.Name()

		if err := stage.Run(ctx, wf); err != nil {
			return fmt.Errorf("%s: %w", stage.Name(), err)
		}

		wf.CompletedStages = append(
			wf.CompletedStages,
			stage.Name(),
		)
	}

	return nil
}
