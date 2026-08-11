package jobs

import (
	"sync"

	"github.com/MichaelVanDerBlond/hashcenter/internal/hashcat"
)

type Job struct {
	SessionID string

	HashFile string

	Dictionary string

	AttackMode int
	HashMode   int

	Rule string

	Mask string

	SessionName string

	Device string

	Workload int

	Extra []string
}

type Runtime struct {
	SessionID string

	Runner *hashcat.Runner

	Status *hashcat.Status

	Result *hashcat.Result
}

type Result struct {
	SessionID string
	Result    *hashcat.Result
	Status    *hashcat.Status
}

type Manager struct {
	mu sync.RWMutex

	jobs map[string]*Runtime
}

var defaultManager = New()

func DefaultManager() *Manager {
	return defaultManager
}

func New() *Manager {
	return &Manager{
		jobs: make(map[string]*Runtime),
	}
}
