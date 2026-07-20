package hashcat

import (
	"context"
	"strings"
)

func (r *Runner) Version(ctx context.Context) (string, error) {

	cmd := r.Command(ctx, "--version")

	out, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(out)), nil
}
