package doctor

type Doctor struct {
	Checks []Check
}

func (d *Doctor) Run() []Result {
	results := make([]Result, 0, len(d.Checks))

	for _, check := range d.Checks {
		results = append(results, check.Run())
	}

	return results
}
