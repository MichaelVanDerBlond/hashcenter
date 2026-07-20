package application

import (
	"github.com/MichaelVanDerBlond/hashcenter/internal/repository"
	"github.com/MichaelVanDerBlond/hashcenter/internal/session"
)

type Application struct {
	Sessions     *session.Manager
	SessionRepo  repository.SessionRepository
	Dictionaries *repository.DictionaryRepository
}

func New(
	sessionRepo repository.SessionRepository,
	dictionaries *repository.DictionaryRepository,
) *Application {
	return &Application{
		Sessions:     session.DefaultManager(),
		SessionRepo:  sessionRepo,
		Dictionaries: dictionaries,
	}
}
