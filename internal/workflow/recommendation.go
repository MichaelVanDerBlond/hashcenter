package workflow

func (p *Plan) Add(step string) {
	p.Steps = append(p.Steps, step)
}
