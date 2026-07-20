package jobs

func Active() []*Runtime {
	return DefaultManager().List()
}
