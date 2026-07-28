package system

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
	"strings"
)

type Result struct {
	Stdout string
	Stderr string
}

func Exec(ctx context.Context, name string, args ...string) (Result, error) {

	cmd := exec.CommandContext(ctx, name, args...)

	var stdout bytes.Buffer
	var stderr bytes.Buffer

	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()

	res := Result{
		Stdout: strings.TrimSpace(stdout.String()),
		Stderr: strings.TrimSpace(stderr.String()),
	}

	if err == nil {
		return res, nil
	}

	if res.Stderr != "" {
		return res, fmt.Errorf("%s: %s: %w", name, res.Stderr, err)
	}

	return res, fmt.Errorf("%s: %w", name, err)
}
