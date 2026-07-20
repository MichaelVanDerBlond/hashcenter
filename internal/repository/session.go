package repository

import "github.com/MichaelVanDerBlond/hashcenter/internal/session"

type SessionRepository interface {
	Save(*session.Session) error

	Update(*session.Session) error

	Get(id string) (*session.Session, error)

	List() ([]session.Session, error)
}
