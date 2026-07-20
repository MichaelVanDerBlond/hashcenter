package session

type State string

const (
	Created  State = "created"
	Running  State = "running"
	Finished State = "finished"
	Failed   State = "failed"
)
