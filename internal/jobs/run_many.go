package jobs

import "context"

type MultiResult struct {
	Results []*Result
	Succeed int
	Failed  int
}

func (m *Manager) RunMany(ctx context.Context, jobs []Job) (*MultiResult, error) {
	out := &MultiResult{
		Results: make([]*Result, 0, len(jobs)),
	}

	for _, job := range jobs {
		result, err := m.Run(ctx, job)
		if err != nil {
			out.Failed++
			continue
		}

		out.Succeed++
		out.Results = append(out.Results, result)
	}

	return out, nil
}
