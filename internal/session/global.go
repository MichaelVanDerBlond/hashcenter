package session

var defaultManager = NewManager()

func DefaultManager() *Manager {
	return defaultManager
}

func ResetDefaultManager() {
	defaultManager = NewManager()
}
