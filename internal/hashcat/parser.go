package hashcat

import "strings"

type Parser struct {
	Status *Status
}

type parserRule struct {
	prefix string
	update func(string)
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

	rules := []parserRule{
		{"Speed.#", p.Status.SetSpeed},
		{"Progress", p.Status.SetProgress},
		{"Recovered", p.Status.SetRecovered},
		{"Time.Estimated", p.Status.SetETA},
		{"Status", p.Status.SetState},
	}

	for _, rule := range rules {
		if strings.HasPrefix(line, rule.prefix) {
			rule.update(line)
			return
		}
	}
}
