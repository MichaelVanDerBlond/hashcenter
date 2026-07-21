package workflow

import "context"

type Stage interface {
	Name() string

	ShouldRun(*Context) bool

	Run(context.Context, *Context) error
}

type StageFunc struct {
	StageName string

	Condition func(*Context) bool

	Handler func(context.Context, *Context) error
}

func (s StageFunc) Name() string {
	return s.StageName
}

func (s StageFunc) ShouldRun(ctx *Context) bool {
	if s.Condition == nil {
		return true
	}

	return s.Condition(ctx)
}

func (s StageFunc) Run(
	ctx context.Context,
	wf *Context,
) error {
	return s.Handler(ctx, wf)
}
