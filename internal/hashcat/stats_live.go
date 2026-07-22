package hashcat

func (p *Parser) Stats() Stats {
	if p == nil || p.Status == nil {
		return Stats{}
	}

	return p.Status.Stats()
}
