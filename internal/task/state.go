package task

type State string

const (
	Queued   State = "queued"
	Running  State = "running"
	Finished State = "finished"
	Failed   State = "failed"
)
