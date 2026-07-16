package analyze

func Analyze(path string) (*Report, error) {

	r := &Report{
		Path: path,
	}

	return r, nil
}
