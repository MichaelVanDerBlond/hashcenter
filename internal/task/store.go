package task

type Store interface {
	Save(*Task) error
	Delete(id string) error
	Load() ([]*Task, error)
}

type MemoryStore struct{}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{}
}

func (s *MemoryStore) Save(*Task) error {
	return nil
}

func (s *MemoryStore) Delete(string) error {
	return nil
}

func (s *MemoryStore) Load() ([]*Task, error) {
	return nil, nil
}
