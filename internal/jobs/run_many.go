package jobs

import (
	"context"
	"errors"
)

type MultiResult struct {
	Results []*Result
	Succeed int
	Failed  int

	Errors []error
}

func (m *Manager) RunMany(ctx context.Context, jobs []Job) (*MultiResult, error) {
	out := &MultiResult{
		Results: make([]*Result, 0, len(jobs)),
		Errors:  make([]error, 0),
	}

	if len(jobs) == 0 {
		return out, nil
	}

	for _, job := range jobs {

		if err := ctx.Err(); err != nil {
			out.Errors = append(out.Errors, err)
			break
		}

		result, err := m.Run(ctx, job)
		if err != nil {
			out.Failed++
			out.Errors = append(out.Errors, err)
			continue
		}

		out.Succeed++
		out.Results = append(out.Results, result)
	}

	if len(out.Errors) > 0 {
		return out, errors.Join(out.Errors...)
	}

	return out, nil
}
