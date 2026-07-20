package hashcat

import "context"

func Available(ctx context.Context) bool {

	r := New()

	_, err := r.Version(ctx)

	return err == nil
}
