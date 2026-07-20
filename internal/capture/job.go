package capture

import "time"

type Job struct {
	Interface string

	Channel int

	Output string

	Duration time.Duration

	Backend string
}
