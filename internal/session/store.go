package session

import (
	"encoding/json"
	"os"
)

const file = "output/session.json"

func Load() (*Session, error) {

	s := &Session{}

	data, err := os.ReadFile(file)
	if err != nil {
		return s, err
	}

	err = json.Unmarshal(data, s)
	return s, err
}

func Save(s *Session) error {

	if err := os.MkdirAll("output", 0755); err != nil {
		return err
	}

	data, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(file, data, 0644)
}
