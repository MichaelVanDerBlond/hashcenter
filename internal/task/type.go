package task

type Type string

const (
	Capture Type = "capture"
	Analyze Type = "analyze"
	Convert Type = "convert"
	Attack  Type = "attack"
	Report  Type = "report"
)
