package workflow

import "context"

type StageFunc struct {
	StageName string
	Handler   func(context.Context, *Context) error
}

func (s StageFunc) Name() string {
	return s.StageName
}

func (s StageFunc) Run(
	ctx context.Context,
	wf *Context,
) error {
	return s.Handler(ctx, wf)
}
