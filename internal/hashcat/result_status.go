package hashcat

func (r *Result) Success() bool {

	if r == nil {
		return false
	}

	return r.ExitCode == 0
}

func (r *Result) Cracked() bool {

	if r == nil {
		return false
	}

	return r.State == "Cracked"
}

func (r *Result) Exhausted() bool {

	if r == nil {
		return false
	}

	return r.State == "Exhausted"
}

func (r *Result) Running() bool {

	if r == nil {
		return false
	}

	return r.State == "Running"
}
