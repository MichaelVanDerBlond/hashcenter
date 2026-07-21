package convert

import "github.com/MichaelVanDerBlond/hashcenter/internal/system"

func Run(path string) (*Result, error) {
	plan, _, err := system.ExecuteConversion(path)
	if err != nil {
		return nil, err
	}

	r := &Result{
		Input:   path,
		Success: true,
	}

	if plan != nil {
		r.Output = plan.OutputFile
		r.Backend = plan.Backend
	}

	return r, nil
}
