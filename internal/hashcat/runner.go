package hashcat

import (
	"context"
	"os/exec"
	"sync"
)

const DefaultBinary = "hashcat"

type Runner struct {
	Binary string

	mu sync.Mutex

	cmd *exec.Cmd

	done chan struct{}
}

func New() *Runner {
	return &Runner{
		Binary: DefaultBinary,
		done:   make(chan struct{}),
	}
}

func (r *Runner) Command(ctx context.Context, args ...string) *exec.Cmd {
	return exec.CommandContext(ctx, r.Binary, args...)
}
