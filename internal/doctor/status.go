package doctor

type Status int

const (
	StatusOK Status = iota
	StatusWarning
	StatusError
)

func (s Status) Icon() string {
	switch s {
	case StatusOK:
		return "✔"
	case StatusWarning:
		return "!"
	case StatusError:
		return "✖"
	default:
		return "?"
	}
}
