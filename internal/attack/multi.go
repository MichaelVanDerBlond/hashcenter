package attack

import (
	"context"

	"github.com/MichaelVanDerBlond/hashcenter/internal/jobs"
)

func ExecuteMany(ctx context.Context, list []Job) (*jobs.MultiResult, error) {
	return jobs.DefaultManager().RunMany(ctx, list)
}
