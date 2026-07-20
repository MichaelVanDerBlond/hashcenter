package system

import (
	"context"
	"os/exec"
)

func Start(ctx context.Context, name string, args ...string) (*exec.Cmd, error) {

	cmd := exec.CommandContext(ctx, name, args...)

	if err := cmd.Start(); err != nil {
		return nil, err
	}

	return cmd, nil
}
