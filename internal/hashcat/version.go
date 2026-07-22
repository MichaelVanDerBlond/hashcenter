package hashcat

import (
	"context"
	"strings"
)

func (r *Runner) Version(ctx context.Context) (string, error) {
	out, err := r.Command(ctx, "--version").Output()
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(out)), nil
}
