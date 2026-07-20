package session

import "github.com/MichaelVanDerBlond/hashcenter/internal/hashcat"

type Runtime struct {
	Session *Session

	Runner *hashcat.Runner

	Status *hashcat.Status
}
