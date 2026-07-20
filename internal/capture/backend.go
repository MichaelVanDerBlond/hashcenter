package capture

import "fmt"

type Backend interface {
	Name() string

	Capture(Job) (Result, error)
}

func Run(b Backend, job Job) (Result, error) {

	if b == nil {
		return Result{}, fmt.Errorf("capture backend is nil")
	}

	return b.Capture(job)
}
