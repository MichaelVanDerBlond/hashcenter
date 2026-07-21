package hashcatruntime

import "sync"

var (
	mu sync.RWMutex

	sessions = make(map[string]*Session)
)

func Register(s *Session) {

	mu.Lock()
	defer mu.Unlock()

	sessions[s.Snapshot().Session] = s
}

func Unregister(name string) {

	mu.Lock()
	defer mu.Unlock()

	delete(sessions, name)
}

func Get(name string) (*Session, bool) {

	mu.RLock()
	defer mu.RUnlock()

	s, ok := sessions[name]

	return s, ok
}

func List() []Snapshot {

	mu.RLock()
	defer mu.RUnlock()

	result := make([]Snapshot, 0, len(sessions))

	for _, s := range sessions {
		result = append(result, s.Snapshot())
	}

	return result
}
