package cli

import (
	"fmt"

	"github.com/MichaelVanDerBlond/hashcenter/internal/database"
	"github.com/MichaelVanDerBlond/hashcenter/internal/repository"
)

func Sessions() error {

	db, err := database.Open()
	if err != nil {
		return err
	}
	defer db.Close()

	repo := repository.NewSQLiteSessionRepository(db)

	sessions, err := repo.List()
	if err != nil {
		return err
	}

	if len(sessions) == 0 {
		fmt.Println("No sessions found.")
		return nil
	}

	fmt.Printf(
		"%-36s %-10s %-19s %-10s %s\n",
		"ID",
		"STATE",
		"STARTED",
		"BACKEND",
		"CAPTURE",
	)

	for _, s := range sessions {

		started := "-"

		if !s.Started.IsZero() {
			started = s.Started.Format("2006-01-02 15:04:05")
		}

		fmt.Printf(
			"%-36s %-10s %-19s %-10s %s\n",
			s.ID,
			s.State,
			started,
			s.Backend,
			s.CaptureFile,
		)
	}

	return nil
}
