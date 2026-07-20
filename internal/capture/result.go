package capture

import "time"

type Result struct {
	Started time.Time

	Finished time.Time

	Duration time.Duration

	Output string

	Success bool
}
