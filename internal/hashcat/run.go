package hashcat

import (
	"bytes"
	"context"
	"time"
)

func (r *Runner) Run(ctx context.Context, args ...string) (*Result, error) {

	cmd := r.Command(ctx, args...)

	var stdout bytes.Buffer
	var stderr bytes.Buffer

	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	result := &Result{
		Started: time.Now(),
	}

	err := cmd.Run()

	result.Finished = time.Now()
	result.Duration = result.Finished.Sub(result.Started)
	result.Stdout = stdout.String()
	result.Stderr = stderr.String()
	result.Error = err

	if cmd.ProcessState != nil {
		result.ExitCode = cmd.ProcessState.ExitCode()
	}

	return result, err
}
