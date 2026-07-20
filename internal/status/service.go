package status

import (
	"github.com/MichaelVanDerBlond/hashcenter/internal/repository"
	"github.com/MichaelVanDerBlond/hashcenter/internal/session"
)

type Service struct {
	Sessions *session.Manager

	Dictionaries *repository.DictionaryRepository
}

func New(dicts *repository.DictionaryRepository) *Service {
	return &Service{
		Sessions:     session.DefaultManager(),
		Dictionaries: dicts,
	}
}

func (s *Service) Snapshot() (*Status, error) {

	st := &Status{
		ActiveSessions: len(s.Sessions.List()),
	}

	if s.Dictionaries == nil {
		return st, nil
	}

	fav, err := s.Dictionaries.FavoriteCount()
	if err != nil {
		return nil, err
	}

	other, err := s.Dictionaries.OtherCount()
	if err != nil {
		return nil, err
	}

	st.FavoriteDictionaries = fav
	st.OtherDictionaries = other

	return st, nil
}
