package hashcat

import "strings"

type Parser struct {
	Status *Status
}

func NewParser() *Parser {
	return &Parser{
		Status: NewStatus(),
	}
}

func (p *Parser) Parse(line string) {

	line = strings.TrimSpace(line)

	if line == "" {
		return
	}

	switch {

	case strings.HasPrefix(line, "Speed.#"):
		p.Status.SetSpeed(line)

	case strings.HasPrefix(line, "Progress"):
		p.Status.SetProgress(line)

	case strings.HasPrefix(line, "Recovered"):
		p.Status.SetRecovered(line)

	case strings.HasPrefix(line, "Time.Estimated"):
		p.Status.SetETA(line)

	case strings.HasPrefix(line, "Status"):
		p.Status.SetState(line)
	}
}
