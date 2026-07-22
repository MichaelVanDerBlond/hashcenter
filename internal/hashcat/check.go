package hashcat

import "context"

func Available(ctx context.Context) bool {
	_, err := New().Version(ctx)
	return err == nil
}
