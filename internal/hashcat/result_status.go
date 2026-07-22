package hashcat

func (r *Result) Success() bool {
	return r != nil && r.ExitCode == 0
}

func (r *Result) Cracked() bool {
	return r != nil && r.State == "Cracked"
}

func (r *Result) Exhausted() bool {
	return r != nil && r.State == "Exhausted"
}

func (r *Result) Running() bool {
	return r != nil && r.State == "Running"
}
