package hashcat

func (r *Result) Success() bool {
	return r != nil && r.ExitCode == 0
}

func (r *Result) HasState(state string) bool {
	return r != nil && r.State == state
}

func (r *Result) Cracked() bool {
	return r.HasState("Cracked")
}

func (r *Result) Exhausted() bool {
	return r.HasState("Exhausted")
}

func (r *Result) Running() bool {
	return r.HasState("Running")
}
