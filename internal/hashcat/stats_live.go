package hashcat

func (p *Parser) Stats() Stats {
	return p.Status.Stats()
}
