package application

func (a *Application) Healthy() bool {
	return a != nil &&
		a.Sessions != nil &&
		a.SessionRepo != nil &&
		a.Dictionaries != nil
}
